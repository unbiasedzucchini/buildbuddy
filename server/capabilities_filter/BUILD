load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "capabilities_filter",
    srcs = ["capabilities_filter.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/capabilities_filter",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:capability_go_proto",
        "//server/environment",
        "//server/util/capabilities",
        "//server/util/claims",
        "//server/util/status",
    ],
)

