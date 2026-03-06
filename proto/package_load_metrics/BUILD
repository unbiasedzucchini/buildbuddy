load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "package_load_metrics",
    srcs = [
        "package_load_metrics.pb.go",
        "package_load_metrics_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/package_load_metrics",
    deps = [
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@com_github_planetscale_vtprotobuf//types/known/durationpb",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/durationpb",
    ],
)

go_test(
    name = "package_load_metrics_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:package_load_metrics.proto",
    ],
    rundir = ".",
)
