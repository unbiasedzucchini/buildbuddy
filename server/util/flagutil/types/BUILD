load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "types",
    srcs = ["types.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/flagutil/types",
    visibility = ["//visibility:public"],
    deps = [
        "//server/util/alert",
        "//server/util/flagutil/common",
        "//server/util/flagutil/types/autoflags/tags",
        "//server/util/log",
        "//server/util/status",
        "@in_gopkg_yaml_v3//:yaml_v3",
    ],
)

