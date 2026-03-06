load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "hit_tracker_client",
    srcs = ["hit_tracker_client.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/hit_tracker_client",
    deps = [
        "//enterprise/server/util/proxy_util",
        "//proto:cache_go_proto",
        "//proto:hit_tracker_go_proto",
        "//proto:remote_execution_go_proto",
        "//proto:resource_go_proto",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/util/alert",
        "//server/util/authutil",
        "//server/util/claims",
        "//server/util/grpc_client",
        "//server/util/log",
        "//server/util/status",
        "//server/util/usageutil",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//types/known/durationpb",
    ],
)

