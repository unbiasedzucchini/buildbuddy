load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "watchdog",
    srcs = ["watchdog.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/watchdog",
    visibility = ["//visibility:public"],
    deps = ["@com_github_jonboulle_clockwork//:clockwork"],
)

