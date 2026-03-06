load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "failure_details",
    srcs = [
        "failure_details.pb.go",
        "failure_details_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/failure_details",
    deps = [
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/descriptorpb",
    ],
)

go_test(
    name = "failure_details_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:failure_details.proto",
    ],
    rundir = ".",
)
