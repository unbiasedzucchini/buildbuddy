load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "composable_cache",
    srcs = ["composable_cache.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/composable_cache",
    deps = [
        "//proto:remote_execution_go_proto",
        "//proto:resource_go_proto",
        "//server/interfaces",
        "//server/remote_cache/digest",
    ],
)

