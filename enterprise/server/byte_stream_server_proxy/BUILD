load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "byte_stream_server_proxy",
    srcs = ["byte_stream_server_proxy.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/byte_stream_server_proxy",
    visibility = ["//visibility:public"],
    deps = [
        "//enterprise/server/remote_crypter",
        "//enterprise/server/util/proxy_util",
        "//proto:remote_execution_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/remote_cache/cachetools",
        "//server/remote_cache/chunking",
        "//server/remote_cache/digest",
        "//server/util/authutil",
        "//server/util/bytebufferpool",
        "//server/util/cdc",
        "//server/util/claims",
        "//server/util/compression",
        "//server/util/lib/set",
        "//server/util/log",
        "//server/util/prefix",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_prometheus_client_golang//prometheus",
        "@io_opentelemetry_go_otel//attribute",
        "@org_golang_google_genproto_googleapis_bytestream//:bytestream",
        "@org_golang_x_sync//errgroup",
    ],
)

