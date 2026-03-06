load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "build_event_stream",
    srcs = [
        "build_event_stream.pb.go",
        "build_event_stream_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/build_event_stream",
    deps = [
        "//proto/action_cache",
        "//proto/command_line",
        "//proto/failure_details",
        "//proto/invocation_policy",
        "//proto/package_load_metrics",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@com_github_planetscale_vtprotobuf//types/known/anypb",
        "@com_github_planetscale_vtprotobuf//types/known/durationpb",
        "@com_github_planetscale_vtprotobuf//types/known/timestamppb",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

go_test(
    name = "build_event_stream_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:build_event_stream.proto",
    ],
    rundir = ".",
)
