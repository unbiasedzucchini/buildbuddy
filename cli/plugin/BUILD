load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//cli:__subpackages__"])

go_library(
    name = "plugin",
    srcs = ["plugin.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/cli/plugin",
    deps = [
        "//cli/arg",
        "//cli/bazelisk",
        "//cli/config",
        "//cli/log",
        "//cli/parser",
        "//cli/storage",
        "//cli/terminal",
        "//cli/workspace",
        "//server/util/disk",
        "//server/util/git",
        "//server/util/shlex",
        "//server/util/status",
        "@com_github_creack_pty//:pty",
        "@in_gopkg_yaml_v2//:yaml_v2",
    ],
)

