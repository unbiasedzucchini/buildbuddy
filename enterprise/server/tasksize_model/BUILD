load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "tasksize_model",
    srcs = ["tasksize_model.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/tasksize_model",
    deps = [
        "//proto:remote_execution_go_proto",
        "//proto:scheduler_go_proto",
        "//server/environment",
        "//server/metrics",
        "//server/util/grpc_client",
        "//server/util/log",
        "//server/util/platform",
        "//server/util/status",
        "@com_github_buildbuddy_io_tensorflow_proto//tensorflow/core/framework",
        "@com_github_buildbuddy_io_tensorflow_proto//tensorflow_serving/apis",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_x_sync//errgroup",
    ],
)

