load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "common",
    srcs = ["common.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/flagutil/common",
    visibility = [
        "//server/config:__subpackages__",
        "//server/util/flag:__subpackages__",
        "//server/util/flagutil:__subpackages__",
    ],
    deps = [
        "//server/util/status",
        "@in_gopkg_yaml_v3//:yaml_v3",
    ],
)

