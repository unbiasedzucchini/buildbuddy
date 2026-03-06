load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "acl",
    srcs = [
        "acl.pb.go",
        "acl_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/acl",
    deps = [
        "//proto/user_id",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "acl_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:acl.proto",
    ],
    rundir = ".",
)
