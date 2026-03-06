load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "chunkstore",
    srcs = ["chunkstore.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/backends/chunkstore",
    visibility = ["//visibility:public"],
    deps = [
        "//server/interfaces",
        "//server/util/background",
        "//server/util/status",
    ],
)

