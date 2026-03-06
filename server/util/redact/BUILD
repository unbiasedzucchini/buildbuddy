load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "redact",
    srcs = ["redact.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/redact",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:build_event_stream_go_proto",
        "//proto:command_line_go_proto",
        "//server/util/flag",
        "//server/util/git",
        "//server/util/status",
        "@com_github_google_shlex//:shlex",
        "@org_golang_google_protobuf//encoding/prototext",
    ],
)

