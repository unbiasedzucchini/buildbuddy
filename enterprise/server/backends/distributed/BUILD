load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "distributed",
    srcs = ["distributed.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/backends/distributed",
    deps = [
        "//enterprise/server/backends/pubsub",
        "//enterprise/server/util/distributed_client",
        "//enterprise/server/util/heartbeat",
        "//enterprise/server/util/redisutil",
        "//proto:distributed_cache_go_proto",
        "//proto:remote_execution_go_proto",
        "//proto:resource_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/remote_cache/digest",
        "//server/resources",
        "//server/util/alert",
        "//server/util/authutil",
        "//server/util/background",
        "//server/util/claims",
        "//server/util/consistent_hash",
        "//server/util/flag",
        "//server/util/ioutil",
        "//server/util/log",
        "//server/util/lru",
        "//server/util/peerset",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
        "@org_golang_x_sync//errgroup",
    ],
)

