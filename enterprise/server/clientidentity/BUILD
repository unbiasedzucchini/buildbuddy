load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "clientidentity",
    srcs = ["clientidentity.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/clientidentity",
    deps = [
        "//server/interfaces",
        "//server/real_environment",
        "//server/util/authutil",
        "//server/util/flag",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_golang_jwt_jwt_v4//:jwt",
        "@com_github_jonboulle_clockwork//:clockwork",
        "@org_golang_google_grpc//metadata",
    ],
)

