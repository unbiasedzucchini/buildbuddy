load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "hit_tracker_service",
    srcs = ["hit_tracker_service.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/hit_tracker_service",
    deps = [
        "//proto:cache_go_proto",
        "//proto:hit_tracker_go_proto",
        "//proto:resource_go_proto",
        "//server/interfaces",
        "//server/real_environment",
        "//server/util/alert",
        "//server/util/log",
        "//server/util/status",
    ],
)

