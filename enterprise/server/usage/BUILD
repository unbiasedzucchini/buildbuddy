load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "usage",
    srcs = ["usage.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/usage",
    deps = [
        "//enterprise/server/usage/config",
        "//enterprise/server/util/redisutil",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/tables",
        "//server/usage/sku",
        "//server/util/alert",
        "//server/util/authutil",
        "//server/util/clickhouse/schema",
        "//server/util/db",
        "//server/util/log",
        "//server/util/platform",
        "//server/util/region",
        "//server/util/status",
        "//server/util/usageutil",
        "@com_github_go_redis_redis_v8//:redis",
        "@com_github_jonboulle_clockwork//:clockwork",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_x_sync//errgroup",
    ],
)

