load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fsutil",
    srcs = ["fsutil.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/fsutil",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:remote_execution_go_proto",
    ],
)

