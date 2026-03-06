load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "action_cache",
    srcs = [
        "action_cache.pb.go",
        "action_cache_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/action_cache",
    deps = [
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "action_cache_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:action_cache.proto",
    ],
    rundir = ".",
)
