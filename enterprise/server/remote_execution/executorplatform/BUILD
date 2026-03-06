load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "executorplatform",
    srcs = ["executorplatform.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/executorplatform",
    deps = [
        "//enterprise/server/util/ci_runner_env",
        "//proto:remote_execution_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/platform",
        "//server/util/status",
        "//server/util/usageutil",
    ],
)

