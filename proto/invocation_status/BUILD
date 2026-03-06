load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "invocation_status",
    srcs = [
        "invocation_status.pb.go",
        "invocation_status_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/invocation_status",
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "invocation_status_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:invocation_status.proto",
    ],
    rundir = ".",
)
