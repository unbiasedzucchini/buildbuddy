load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "context",
    srcs = [
        "context.pb.go",
        "context_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/context",
    deps = [
        "//proto/user_id",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "context_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:context.proto",
    ],
    rundir = ".",
)
