load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "stat_filter",
    srcs = [
        "stat_filter.pb.go",
        "stat_filter_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/stat_filter",
    deps = [
        "//proto/invocation_status",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)

go_test(
    name = "stat_filter_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:stat_filter.proto",
    ],
    rundir = ".",
)
