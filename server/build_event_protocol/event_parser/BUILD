load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "event_parser",
    srcs = ["event_parser.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/build_event_protocol/event_parser",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:build_event_stream_go_proto",
        "//proto:command_line_go_proto",
        "//proto:invocation_go_proto",
        "//server/build_event_protocol/invocation_format",
        "//server/util/git",
        "//server/util/timeutil",
    ],
)

