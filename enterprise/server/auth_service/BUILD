load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "auth_service",
    srcs = ["auth_service.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/auth_service",
    deps = [
        "//proto:auth_go_proto",
        "//server/interfaces",
        "//server/real_environment",
        "//server/util/authutil",
        "//server/util/claims",
        "//server/util/status",
        "//server/util/subdomain",
        "@com_github_golang_jwt_jwt_v4//:jwt",
    ],
)

