load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "user_id",
    srcs = [
        "user_id.pb.go",
        "user_id_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/user_id",
    deps = [
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
    ],
)

go_test(
    name = "user_id_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:user_id.proto",
    ],
    rundir = ".",
)
