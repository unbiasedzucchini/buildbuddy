load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "common",
    srcs = [
        "common.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/api/v1/common",
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

go_test(
    name = "common_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto/api/v1:common.proto",
    ],
    rundir = ".",
)
