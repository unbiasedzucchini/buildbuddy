load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "copy_on_write",
    srcs = ["copy_on_write.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/copy_on_write",
    visibility = ["//visibility:public"],
    deps = [
        "//enterprise/server/remote_execution/snaputil",
        "//proto:remote_execution_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/remote_cache/digest",
        "//server/resources",
        "//server/util/alert",
        "//server/util/boundedstack",
        "//server/util/flag",
        "//server/util/lockmap",
        "//server/util/log",
        "//server/util/lru",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_x_sync//errgroup",
        "@org_golang_x_sys//unix",
        "@org_golang_x_time//rate",
    ],
)

