load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "filter",
    srcs = ["filter.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/filter",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:invocation_status_go_proto",
        "//proto:stat_filter_go_proto",
        "//server/util/query_builder",
        "//server/util/status",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
    ],
)

