load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gossip",
    srcs = ["gossip.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/gossip",
    visibility = ["//visibility:public"],
    deps = [
        "//server/interfaces",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/network",
        "//server/util/retry",
        "//server/util/statusz",
        "@com_github_hashicorp_memberlist//:memberlist",
        "@com_github_hashicorp_serf//serf",
        "@com_github_rs_zerolog//:zerolog",
    ],
)

