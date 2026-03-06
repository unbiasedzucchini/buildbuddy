load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "commandutil",
    srcs = [
        "commandutil.go",
        "commandutil_unix.go",
        "commandutil_windows.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/commandutil",
    deps = [
        "//enterprise/server/util/procstats",
        "//proto:execution_stats_go_proto",
        "//proto:remote_execution_go_proto",
        "//server/interfaces",
        "//server/util/log",
        "//server/util/status",
        "@com_github_mitchellh_go_ps//:go-ps",
    ] + select({
        "@io_bazel_rules_go//go/platform:windows": [
            "@com_github_shirou_gopsutil_v3//process",
            "@org_golang_x_sys//windows",
        ],
        "//conditions:default": [],
    }),
)

