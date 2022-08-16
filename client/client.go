package main

import (
	"context"
	"fmt"
	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	httpreport "github.com/openzipkin/zipkin-go/reporter/http"
	"go_test_gin/zipkin/server/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)


const (
	serviceName        = "simple_zipkin_server"
	zipkinAddr         = "http://10.100.48.199:9411/api/v2/spans"
	zipkinRecorderAddr = "10.100.48.199:9000"
)

func main()  {

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

	//_, err = grpc.Dial("127.0.0.1:50052", grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)) )
	//defer conn.Close()

	clientConn, err := grpc.Dial("127.0.0.1:50052", grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)),grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}

	helloServiceClient := helloworld.NewHellowClient(clientConn)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)

	helloResponse, err := helloServiceClient.HelloWord(ctx, &helloworld.HelloReq{
		Name: "hahhahhaa",
	})

	cancelFunc()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(helloResponse)
}
