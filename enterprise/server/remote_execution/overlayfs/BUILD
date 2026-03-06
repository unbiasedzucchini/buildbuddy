load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "overlayfs",
    srcs = [
        "overlayfs.go",
        "overlayfs_linux.go",
        "overlayfs_notlinux.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/overlayfs",
    visibility = ["//visibility:public"],
    deps = select({
        "@io_bazel_rules_go//go/platform:darwin": [
            "//server/util/status",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "//server/util/disk",
            "//server/util/status",
            "//server/util/tracing",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "//server/util/status",
        ],
        "//conditions:default": [],
    }),
)

