package main

import (
	"fmt"

	pb "github.com/misterwilliam/bazel_grpc/simple"
)

func main() {
	fmt.Println(&pb.Foo{Blob: "hi"})
}
