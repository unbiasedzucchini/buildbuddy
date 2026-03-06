load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "retry",
    srcs = ["retry.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/retry",
    visibility = ["//visibility:public"],
    deps = [
        "//server/util/log",
        "@com_github_jonboulle_clockwork//:clockwork",
    ],
)

