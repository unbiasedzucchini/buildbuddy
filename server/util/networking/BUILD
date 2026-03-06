load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("//dockerfiles/test_images:defs.bzl", "NET_TOOLS_IMAGE")

go_library(
    name = "networking",
    srcs = ["networking.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/networking",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:remote_execution_go_proto",
        "//server/metrics",
        "//server/util/alert",
        "//server/util/background",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/random",
        "//server/util/status",
        "//server/util/tracing",
        "//server/util/uuid",
        "@com_github_prometheus_client_golang//prometheus",
        "@com_github_vishvananda_netlink//:netlink",
        "@org_golang_x_sync//errgroup",
        "@org_golang_x_sys//unix",
    ],
)

