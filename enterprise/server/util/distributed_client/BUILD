load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "distributed_client",
    srcs = ["distributed_client.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/util/distributed_client",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:distributed_cache_go_proto",
        "//proto:remote_execution_go_proto",
        "//proto:resource_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/remote_cache/config",
        "//server/remote_cache/digest",
        "//server/resources",
        "//server/util/authutil",
        "//server/util/bytebufferpool",
        "//server/util/grpc_client",
        "//server/util/grpc_server",
        "//server/util/ioutil",
        "//server/util/log",
        "//server/util/prefix",
        "//server/util/status",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//reflection",
    ],
)

