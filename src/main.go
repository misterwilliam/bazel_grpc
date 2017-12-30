package main

import (
	"fmt"
	"log"
	"net"

	"lib"
	"no_go_package"

  _ "github.com/misterwilliam/bazel_grpc/parent"
	pb "github.com/misterwilliam/bazel_grpc/simple"
	pb2 "github.com/misterwilliam/bazel_grpc/src/has_external_deps"
	pb3 "github.com/misterwilliam/bazel_grpc/src/has_external_deps_grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetFoo(ctx context.Context, req *pb.GetReq) (*pb.Foo, error) {
	return &pb.Foo{}, nil
}

func (s *server) Annotated(ctx context.Context, req *pb3.Foo) (*pb3.Foo, error) {
	return &pb3.Foo{}, nil
}

func main() {
	fmt.Println(&pb.Foo{Blob: "hi"})
	fmt.Println(&pb2.Foo{Very: "hi"})

	fmt.Println(no_go_package.Foo{Blob: "yes"})
	fmt.Println(lib.Foo{Blob: "yes"})
	//fmt.Println(lib.Bar{Blob: "yes"})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Listen failed")
	}
	srv := &server{}
	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, srv)
	pb3.RegisterHasAnnotationServer(s, srv)
	s.Serve(lis)
}
