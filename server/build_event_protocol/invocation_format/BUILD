load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "invocation_format",
    srcs = ["invocation_format.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/build_event_protocol/invocation_format",
    visibility = ["//visibility:public"],
    deps = [
        "//proto/invocation",
        "//server/util/status",
    ],
)

