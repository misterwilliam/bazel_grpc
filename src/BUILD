load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "main",
    srcs = ["main.go"],
    importpath = "github.com/misterwilliam/bazel_grpc/main",
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
