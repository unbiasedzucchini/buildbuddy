load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "rangecache",
    srcs = ["rangecache.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/raft/rangecache",
    deps = [
        "//proto:raft_go_proto",
        "//server/metrics",
        "//server/util/log",
        "//server/util/rangemap",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_prometheus_client_golang//prometheus",
        "@io_opentelemetry_go_otel//attribute",
    ],
)

