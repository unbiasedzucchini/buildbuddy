load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "userdb",
    srcs = ["userdb.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/backends/userdb",
    deps = [
        "//enterprise/server/util/keystore",
        "//proto:capability_go_proto",
        "//proto:group_go_proto",
        "//proto:telemetry_go_proto",
        "//proto:user_id_go_proto",
        "//proto:user_list_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/tables",
        "//server/util/authutil",
        "//server/util/capabilities",
        "//server/util/claims",
        "//server/util/db",
        "//server/util/log",
        "//server/util/query_builder",
        "//server/util/role",
        "//server/util/status",
        "//server/util/subdomain",
    ],
)

