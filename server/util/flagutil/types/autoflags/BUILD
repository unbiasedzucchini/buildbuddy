load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "autoflags",
    srcs = ["autoflags.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/flagutil/types/autoflags",
    visibility = [
        "//server/util/flag:__subpackages__",
        "//server/util/flagutil:__subpackages__",
    ],
    deps = [
        "//server/util/flagutil/types",
        "//server/util/flagutil/types/autoflags/tags",
        "//server/util/log",
    ],
)

