load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "server_notification",
    srcs = ["server_notification.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/server_notification",
    visibility = ["//visibility:public"],
    deps = [
        "//enterprise/server/backends/pubsub",
        "//proto:server_notification_go_proto",
        "//server/real_environment",
        "//server/util/alert",
        "//server/util/proto",
        "//server/util/status",
        "@com_github_go_redis_redis_v8//:redis",
        "@org_golang_google_protobuf//reflect/protoreflect",
    ],
)

