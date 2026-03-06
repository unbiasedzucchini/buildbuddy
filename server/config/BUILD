load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "config",
    srcs = ["config.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/config",
    visibility = ["//visibility:public"],
    deps = [
        "//server/interfaces",
        "//server/util/flag",
        "//server/util/flagutil",
        "//server/util/flagutil/common",
        "//server/util/flagutil/yaml",
        "//server/util/log",
        "//server/util/status",
    ],
)

