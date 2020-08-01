package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/janowsiany/grpc-web-example/protobuf"
)

//go:generate go get github.com/golang/protobuf/protoc-gen-go
//go:generate protoc --proto_path=protobuf/proto/ --go_out=plugins=grpc:protobuf protobuf/proto/service.proto

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &greeterService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type greeterService struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return nil, status.Error(codes.InvalidArgument, "auth.invalid-username")
}
