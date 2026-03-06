load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//cli:__subpackages__"])

go_library(
    name = "help",
    srcs = ["help.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/cli/help",
    deps = [
        "//cli/bazelisk",
        "//cli/cli_command",
        "//cli/parser/arguments",
        "//cli/parser/parsed",
        "//cli/version",
    ],
)

