load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "capability",
    srcs = [
        "capability.pb.go",
        "capability_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/capability",
    deps = [
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "capability_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:capability.proto",
    ],
    rundir = ".",
)
