package main

import (
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"github.com/openzipkin/zipkin-go/reporter"
	httpreport "github.com/openzipkin/zipkin-go/reporter/http"
	"go_test_gin/zipkin/server/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	serviceName        = "simple_zipkin_server"
	zipkinAddr         = "http://10.100.48.199:9411/api/v2/spans"
	zipkinRecorderAddr = "10.100.48.199:9000"
)
func main()  {
	//启动一个grpc服务
	grpcServer()
}


func grpcServer()  {

	reporter := httpreport.NewReporter(zipkinAddr)
	//reporter := logreporter.NewReporter(log.New(os.Stderr, "", log.LstdFlags))
	defer reporter.Close()
	endpoint, err := zipkin.NewEndpoint(serviceName, zipkinRecorderAddr)
	if err != nil {
		log.Fatalf("unable to create local endpoint: %+v\n", err)
	}
	// initialize our tracer
	tracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		log.Fatalf("unable to create tracer: %+v\n", err)
	}


	li, err := net.Listen("tcp", "127.0.0.1:50052")

	server := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)))

	//注册grpc服务的实现
	hs := &HelloService{}
	helloworld.RegisterHellowServer(server, hs)

	err = server.Serve(li)
	if err != nil {
		panic(err)
	}
}


func NewZipkinTracer(url,serviceName,hostPort string) (*zipkin.Tracer,reporter.Reporter,error) {

	// 初始化zipkin reporter
	// reporter可以有很多种，如：logReporter、httpReporter，这里我们只使用httpReporter将span报告给http服务，也就是zipkin的http后台
	r := httpreport.NewReporter(url)

	//创建一个endpoint，用来标识当前服务，服务名：服务地址和端口
	endpoint, err := zipkin.NewEndpoint(serviceName,hostPort)
	if err != nil {
		return nil,r,err
	}

	// 初始化追踪器 主要作用有解析span，解析上下文等
	tracer, err := zipkin.NewTracer(r,zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		return nil,r,err
	}

	return tracer,r,nil
}

