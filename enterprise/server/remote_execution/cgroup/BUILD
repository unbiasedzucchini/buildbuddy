load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "cgroup",
    srcs = ["cgroup.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/cgroup",
    target_compatible_with = ["@platforms//os:linux"],
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:linux": [
            "//enterprise/server/remote_execution/block_io",
            "//enterprise/server/util/cpuset",
            "//proto:remote_execution_go_proto",
            "//proto:scheduler_go_proto",
            "//server/util/log",
            "//server/util/status",
        ],
        "//conditions:default": [],
    }),
)

