load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "action_cache_server",
    srcs = ["action_cache_server.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/remote_cache/action_cache_server",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:capability_go_proto",
        "//proto:remote_execution_go_proto",
        "//proto:resource_go_proto",
        "//server/environment",
        "//server/hostid",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/remote_cache/chunking",
        "//server/remote_cache/digest",
        "//server/util/bazel_request",
        "//server/util/capabilities",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/prefix",
        "//server/util/proto",
        "//server/util/status",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_x_sync//errgroup",
    ],
)

