load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "cache",
    srcs = [
        "cache.pb.go",
        "cache_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/cache",
    deps = [
        "//proto/context",
        "//proto/remote_execution",
        "//proto/resource",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@com_github_planetscale_vtprotobuf//types/known/durationpb",
        "@com_github_planetscale_vtprotobuf//types/known/fieldmaskpb",
        "@com_github_planetscale_vtprotobuf//types/known/timestamppb",
        "@org_golang_google_genproto_googleapis_rpc//status",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/fieldmaskpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

go_test(
    name = "cache_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:cache.proto",
    ],
    rundir = ".",
)
