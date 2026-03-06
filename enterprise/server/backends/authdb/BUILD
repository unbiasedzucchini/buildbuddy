load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "authdb",
    srcs = ["authdb.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/backends/authdb",
    deps = [
        "//proto:capability_go_proto",
        "//proto:group_go_proto",
        "//proto:user_id_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/tables",
        "//server/util/authutil",
        "//server/util/capabilities",
        "//server/util/claims",
        "//server/util/db",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/lru",
        "//server/util/perms",
        "//server/util/query_builder",
        "//server/util/random",
        "//server/util/role",
        "//server/util/status",
        "//server/util/subdomain",
        "//third_party/singleflight",
        "@com_github_jonboulle_clockwork//:clockwork",
        "@com_github_prometheus_client_golang//prometheus",
        "@io_gorm_gorm//:gorm",
        "@io_gorm_gorm//clause",
        "@org_golang_x_crypto//chacha20",
    ],
)

