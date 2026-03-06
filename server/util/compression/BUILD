load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "compression",
    srcs = ["compression.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/compression",
    visibility = ["//visibility:public"],
    deps = [
        "//server/interfaces",
        "//server/metrics",
        "//server/util/bytebufferpool",
        "//server/util/ioutil",
        "//server/util/log",
        "@com_github_klauspost_compress//zstd",
        "@com_github_prometheus_client_golang//prometheus",
    ],
)

