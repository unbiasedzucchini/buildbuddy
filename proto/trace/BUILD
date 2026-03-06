load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "trace",
    srcs = [
        "trace.pb.go",
        "trace_grpc.pb.go",
        "trace_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/trace",
    deps = [
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "trace_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:trace.proto",
    ],
    rundir = ".",
)
