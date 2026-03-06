load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "content_addressable_storage_server_proxy",
    srcs = ["content_addressable_storage_server_proxy.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/content_addressable_storage_server_proxy",
    visibility = ["//visibility:public"],
    deps = [
        "//enterprise/server/remote_crypter",
        "//enterprise/server/util/proxy_util",
        "//proto:remote_execution_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/remote_cache/digest",
        "//server/util/authutil",
        "//server/util/log",
        "//server/util/proto",
        "//server/util/rpcutil",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_google_grpc//codes",
    ],
)

