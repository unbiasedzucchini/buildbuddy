load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "cache_server",
    srcs = ["cache_server.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/cache_server",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:cache_go_proto",
        "//proto:cache_service_go_proto",
        "//server/environment",
        "//server/real_environment",
        "//server/util/prefix",
        "//server/util/status",
    ],
)

