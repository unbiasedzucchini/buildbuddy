load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "quota",
    srcs = ["quota_manager.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/quota",
    deps = [
        "//proto:group_go_proto",
        "//proto:server_notification_go_proto",
        "//server/environment",
        "//server/metrics",
        "//server/real_environment",
        "//server/util/alert",
        "//server/util/claims",
        "//server/util/log",
        "//server/util/quota",
        "//server/util/status",
        "@com_github_prometheus_client_golang//prometheus",
        "@com_github_throttled_throttled_v2//:throttled",
        "@com_github_throttled_throttled_v2//store/goredisstore.v8:goredisstore_v8",
    ],
)

