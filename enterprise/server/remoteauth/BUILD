load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "remoteauth",
    srcs = ["remoteauth.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remoteauth",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:auth_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/real_environment",
        "//server/tables",
        "//server/util/authutil",
        "//server/util/claims",
        "//server/util/grpc_client",
        "//server/util/log",
        "//server/util/lru",
        "//server/util/status",
        "//server/util/subdomain",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//metadata",
    ],
)

