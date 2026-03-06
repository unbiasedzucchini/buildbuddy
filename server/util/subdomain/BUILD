load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "subdomain",
    srcs = ["subdomain.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/subdomain",
    visibility = ["//visibility:public"],
    deps = [
        "//server/endpoint_urls/build_buddy_url",
        "//server/environment",
        "//server/tables",
        "//server/util/alert",
        "//server/util/flag",
        "//server/util/urlutil",
    ],
)

