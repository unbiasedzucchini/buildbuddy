load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "paging",
    srcs = ["paging.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/paging",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:pagination_go_proto",
        "//server/util/proto",
        "//server/util/status",
    ],
)

