load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "scim",
    srcs = ["scim.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/scim",
    deps = [
        "//enterprise/server/saml",
        "//proto:capability_go_proto",
        "//proto:group_go_proto",
        "//proto:user_id_go_proto",
        "//proto:user_list_go_proto",
        "//server/environment",
        "//server/http/interceptors",
        "//server/interfaces",
        "//server/real_environment",
        "//server/tables",
        "//server/util/alert",
        "//server/util/log",
        "//server/util/role",
        "//server/util/status",
    ],
)

