package main

import (
	"context"
	"fmt"
	helloworld "go_test_gin/zipkin/server/helloworld"
)

type HelloService struct {
	helloworld.UnimplementedHellowServer
}

func (s HelloService) HelloWord(ctx context.Context, request *helloworld.HelloReq ) (*helloworld.HelloResp, error) {
	resp := helloworld.HelloResp{}
	resp.Body = fmt.Sprintf("my name is hello grpc %s", request.Name )
	return &resp, nil
}