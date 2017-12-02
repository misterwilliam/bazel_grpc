package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/misterwilliam/bazel_grpc/simple"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetFoo(ctx context.Context, req *pb.GetReq) (*pb.Foo, error) {
	return &pb.Foo{}, nil
}

func main() {
	fmt.Println(&pb.Foo{Blob: "hi"})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Listen failed")
	}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &server{})
	s.Serve(lis)
}
