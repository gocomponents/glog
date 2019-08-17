package main

import (
	"context"
	"fmt"
	."glog/glogstash/produce_consume"
	"glog/proto"
	"google.golang.org/grpc"
	"net"
)

func main()  {
	go Consume()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterLogStashServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

type server struct{}

func (s *server) Send(ctx context.Context, request *proto.Log) (*proto.Response, error) {
	defer func() {
		if info := recover(); info != nil {
			fmt.Println("post to mq panic", info)
		}
	}()

	go Produce(request)

	return &proto.Response{
		ErrorCode: 0,
		Message:   "",
	}, nil
}
