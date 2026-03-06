load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//cli:__subpackages__"])

go_library(
    name = "versioncmd",
    srcs = ["versioncmd.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/cli/versioncmd",
    deps = [
        "//cli/arg",
        "//cli/plugin",
        "//cli/setup",
        "//cli/version",
        "//server/util/status",
    ],
)

