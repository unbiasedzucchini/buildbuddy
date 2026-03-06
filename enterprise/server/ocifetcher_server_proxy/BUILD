load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "ocifetcher_server_proxy",
    srcs = ["ocifetcher_server_proxy.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/ocifetcher_server_proxy",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:oci_fetcher_go_proto",
        "//server/environment",
        "//server/real_environment",
        "//server/util/status",
    ],
)

