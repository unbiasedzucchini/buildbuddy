load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "action_cache_server_proxy",
    srcs = ["action_cache_server_proxy.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/action_cache_server_proxy",
    visibility = ["//visibility:public"],
    deps = [
        "//enterprise/server/remote_crypter",
        "//enterprise/server/util/proxy_util",
        "//proto:remote_execution_go_proto",
        "//proto:resource_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/remote_cache/cachetools",
        "//server/remote_cache/digest",
        "//server/util/authutil",
        "//server/util/flag",
        "//server/util/prefix",
        "//server/util/proto",
        "//server/util/status",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_google_grpc//status",
    ],
)

