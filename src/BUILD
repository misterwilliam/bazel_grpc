load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "simple_proto",
    srcs = ["simple.proto"],
)

go_proto_library(
    name = "simple_go_proto",
    importpath = "github.com/misterwilliam/bazel_grpc/simple",
    proto = ":simple_proto",
)

go_binary(
    name = "main",
    srcs = ["main.go"],
    importpath = "github.com/misterwilliam/bazel_grpc/main",
    deps = [
        ":simple_go_proto",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/example/project/src",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "src",
    embed = [":go_default_library"],
    importpath = "github.com/example/project/src",
    visibility = ["//visibility:public"],
)
