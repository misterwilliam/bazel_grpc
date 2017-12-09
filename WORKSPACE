# Go rules
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    commit = "70fe781d940c7b9cfee164d421ea85451322ba50",
)
load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
go_rules_dependencies()
go_register_toolchains()

# Go proto and gRPC rules
load("@io_bazel_rules_go//proto:def.bzl", "proto_register_toolchains")
proto_register_toolchains()

GOOGLEAPIS_SHA = "5c6df0cd18c6a429eab739fb711c27f6e1393366" # May 14, 2017
new_http_archive(
        name = "googleapis",
        strip_prefix = "googleapis-" + GOOGLEAPIS_SHA,
        url = "https://github.com/googleapis/googleapis/archive/" + GOOGLEAPIS_SHA + ".tar.gz",
        build_file_content = """
load("@com_google_protobuf//:protobuf.bzl", "cc_proto_library", "py_proto_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

filegroup(
    name = "http_api_protos_src",
    srcs = [
        "google/api/annotations.proto",
        "google/api/http.proto",
    ],
    visibility = ["//visibility:public"],
 )

proto_library(
    name = "http_api_protos_proto",
    srcs = [":http_api_protos_src"],
    deps = ["@com_google_protobuf//:descriptor_proto"],
    visibility = ["//visibility:public"],
)

go_proto_library(
    name = "descriptor_go_proto",
    importpath = "github.com/golang/protobuf/protoc-gen-go/descriptor",
    proto = "@com_google_protobuf//:descriptor_proto",
    visibility = ["//visibility:public"],
)

go_proto_library(
  name = "http_api_go_proto",
  importpath = "google.golang.org/genproto/googleapis/api/annotations",
  visibility = ["//visibility:public"],
  proto = ":http_api_protos_proto",
  deps = [
    "@com_github_golang_protobuf//ptypes/duration:go_default_library",
    ":descriptor_go_proto",
  ],
)
"""
)
