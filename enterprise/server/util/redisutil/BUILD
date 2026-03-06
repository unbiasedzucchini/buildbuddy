load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "redisutil",
    srcs = ["redisutil.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/util/redisutil",
    deps = [
        "//enterprise/server/experiments",
        "//server/interfaces",
        "//server/metrics",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/random",
        "//server/util/retry",
        "//server/util/status",
        "@com_github_cespare_xxhash_v2//:xxhash",
        "@com_github_dgryski_go_rendezvous//:go-rendezvous",
        "@com_github_go_redis_redis_extra_redisotel_v8//:redisotel",
        "@com_github_go_redis_redis_v8//:redis",
    ],
)

