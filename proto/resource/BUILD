load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "resource",
    srcs = [
        "resource.pb.go",
        "resource_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/resource",
    deps = [
        "//proto/remote_execution",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "resource_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:resource.proto",
    ],
    rundir = ".",
)
