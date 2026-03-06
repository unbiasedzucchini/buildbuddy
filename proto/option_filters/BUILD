load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "option_filters",
    srcs = [
        "option_filters.pb.go",
        "option_filters_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/options",
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "option_filters_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:option_filters.proto",
    ],
    rundir = ".",
)
