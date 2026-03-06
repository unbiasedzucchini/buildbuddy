load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "config",
    srcs = ["config.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/workflow/config",
    deps = [
        "//enterprise/server/webhooks/webhook_data",
        "//proto:runner_go_proto",
        "//server/build_event_protocol/accumulator",
        "@in_gopkg_yaml_v2//:yaml_v2",
    ],
)

