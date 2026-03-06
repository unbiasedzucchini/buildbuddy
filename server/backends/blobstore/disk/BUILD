load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "disk",
    srcs = ["disk.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/backends/blobstore/disk",
    visibility = ["//visibility:public"],
    deps = [
        "//server/backends/blobstore/util",
        "//server/interfaces",
        "//server/util/disk",
        "//server/util/ioutil",
        "//server/util/log",
        "//server/util/tracing",
    ],
)

