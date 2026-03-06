load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "target",
    srcs = [
        "target.pb.go",
        "target_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/target",
    deps = [
        "//proto/api/v1/common",
        "//proto/build_event_stream",
        "//proto/context",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@com_github_planetscale_vtprotobuf//types/known/timestamppb",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

go_test(
    name = "target_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:target.proto",
    ],
    rundir = ".",
)
