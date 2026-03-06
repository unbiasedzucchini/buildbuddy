load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//cli:__subpackages__"])

go_library(
    name = "flaghistory",
    srcs = ["flaghistory.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/cli/flaghistory",
    deps = [
        "//cli/arg",
        "//cli/log",
        "//cli/storage",
        "//cli/workspace",
        "//server/util/hash",
        "//server/util/uuid",
    ],
)

