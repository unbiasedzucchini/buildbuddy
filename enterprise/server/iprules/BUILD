load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "iprules",
    srcs = ["iprules.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/iprules",
    deps = [
        "//proto:iprules_go_proto",
        "//proto:server_notification_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/tables",
        "//server/util/alert",
        "//server/util/authutil",
        "//server/util/clientip",
        "//server/util/db",
        "//server/util/log",
        "//server/util/lru",
        "//server/util/status",
        "@com_github_prometheus_client_golang//prometheus",
    ],
)

