load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "clickhouse_backup_lib",
    srcs = ["clickhouse_backup.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/tools/clickhouse_backup",
    visibility = ["//visibility:private"],
    deps = [
        "//server/backends/blobstore",
        "//server/config",
        "//server/environment",
        "//server/real_environment",
        "//server/util/clickhouse",
        "//server/util/flag",
        "//server/util/flagutil",
        "//server/util/log",
    ],
)

go_binary(
    name = "clickhouse_backup",
    embed = [":clickhouse_backup_lib"],
    pure = "on",
    static = "on",
    visibility = ["//visibility:public"],
)

