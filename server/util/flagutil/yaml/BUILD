load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "yaml",
    srcs = ["yaml.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/flagutil/yaml",
    visibility = ["//visibility:public"],
    deps = [
        "//server/util/alert",
        "//server/util/flagutil/common",
        "//server/util/log",
        "//server/util/status",
        "@in_gopkg_yaml_v3//:yaml_v3",
    ],
)

