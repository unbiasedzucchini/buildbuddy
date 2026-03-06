load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "shlex",
    srcs = ["shlex.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/shlex",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_google_shlex//:shlex",
    ],
)

