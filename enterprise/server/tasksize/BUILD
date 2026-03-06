load("@io_bazel_rules_go//go:def.bzl", "go_library")

# gazelle:default_visibility //enterprise:__subpackages__,@buildbuddy_internal//:__subpackages__
package(default_visibility = [
    "//enterprise:__subpackages__",
    "@buildbuddy_internal//:__subpackages__",
])

go_library(
    name = "tasksize",
    srcs = ["tasksize.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/tasksize",
    deps = [
        "//enterprise/server/experiments",
        "//enterprise/server/tasksize_model",
        "//proto:remote_execution_go_proto",
        "//proto:scheduler_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/util/authutil",
        "//server/util/log",
        "//server/util/platform",
        "//server/util/proto",
        "//server/util/status",
        "@com_github_go_redis_redis_v8//:redis",
        "@com_github_prometheus_client_golang//prometheus",
    ],
)

