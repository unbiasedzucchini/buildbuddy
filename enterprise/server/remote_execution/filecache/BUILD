load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = [
    "//enterprise:__subpackages__",
    # For tests only.
    "//server/cache/dirtools:__pkg__",  # keep
])

go_library(
    name = "filecache",
    srcs = ["filecache.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/filecache",
    deps = [
        "//proto:remote_execution_go_proto",
        "//server/interfaces",
        "//server/metrics",
        "//server/remote_cache/digest",
        "//server/util/claims",
        "//server/util/disk",
        "//server/util/fastcopy",
        "//server/util/log",
        "//server/util/lru",
        "//server/util/random",
        "//server/util/status",
        "@com_github_prometheus_client_golang//prometheus",
    ],
)

