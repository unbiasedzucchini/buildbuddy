load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//cli:__subpackages__"])

go_library(
    name = "js",
    srcs = ["js.go"],
    embedsrcs = ["translate.js"],
    importpath = "github.com/buildbuddy-io/buildbuddy/cli/translate/js",
    visibility = ["//visibility:public"],
    deps = [
        "//cli/translate/builtins",
        "@com_github_dop251_goja//:goja",
    ],
)

