load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "interceptors",
    srcs = ["interceptors.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/http/interceptors",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:context_go_proto",
        "//server/capabilities_filter",
        "//server/environment",
        "//server/features",
        "//server/http/csp",
        "//server/http/protolet",
        "//server/metrics",
        "//server/util/alert",
        "//server/util/clientip",
        "//server/util/compression",
        "//server/util/log",
        "//server/util/proto",
        "//server/util/region",
        "//server/util/request_context",
        "//server/util/subdomain",
        "//server/util/uuid",
        "@com_github_prometheus_client_golang//prometheus",
    ],
)

