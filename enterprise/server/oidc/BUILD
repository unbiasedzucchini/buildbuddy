load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "oidc",
    srcs = ["oidc.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/oidc",
    deps = [
        "//enterprise/server/selfauth",
        "//server/endpoint_urls/build_buddy_url",
        "//server/environment",
        "//server/interfaces",
        "//server/tables",
        "//server/util/alert",
        "//server/util/authutil",
        "//server/util/claims",
        "//server/util/cookie",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/random",
        "//server/util/status",
        "//server/util/subdomain",
        "//server/util/urlutil",
        "@com_github_coreos_go_oidc_v3//oidc",
        "@com_github_google_uuid//:uuid",
        "@org_golang_google_grpc//credentials",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//peer",
        "@org_golang_x_oauth2//:oauth2",
    ],
)

