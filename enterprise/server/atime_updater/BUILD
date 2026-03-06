load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "atime_updater",
    srcs = ["atime_updater.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/atime_updater",
    visibility = ["//visibility:public"],
    deps = [
        "//enterprise/server/batch_operator",
        "//proto:remote_execution_go_proto",
        "//server/metrics",
        "//server/real_environment",
        "//server/util/status",
        "@org_golang_google_grpc//status",
    ],
)

