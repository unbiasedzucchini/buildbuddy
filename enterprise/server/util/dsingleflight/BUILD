load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "dsingleflight",
    srcs = ["dsingleflight.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/util/dsingleflight",
    deps = [
        "//server/real_environment",
        "//server/util/log",
        "//server/util/proto",
        "//server/util/status",
        "@com_github_go_redis_redis_v8//:redis",
        "@org_golang_google_genproto_googleapis_rpc//status",
        "@org_golang_google_grpc//status",
    ],
)

