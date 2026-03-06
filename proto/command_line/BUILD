load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "command_line",
    srcs = [
        "command_line.pb.go",
        "command_line_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/command_line",
    deps = [
        "//proto/option_filters",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "command_line_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:command_line.proto",
    ],
    rundir = ".",
)
