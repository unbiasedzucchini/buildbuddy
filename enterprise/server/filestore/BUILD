load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "filestore",
    srcs = ["filestore.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/filestore",
    deps = [
        "//proto:remote_execution_go_proto",
        "//proto:resource_go_proto",
        "//proto:storage_go_proto",
        "//server/interfaces",
        "//server/remote_cache/digest",
        "//server/util/disk",
        "//server/util/log",
        "//server/util/random",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_jonboulle_clockwork//:clockwork",
    ],
)

