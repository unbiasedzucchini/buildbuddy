load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "crypter_service",
    srcs = ["crypter_service.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/crypter_service",
    deps = [
        "//enterprise/server/crypter",
        "//enterprise/server/crypter_key_cache",
        "//proto:encryption_go_proto",
        "//proto:remote_execution_go_proto",
        "//proto:storage_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/tables",
        "//server/util/db",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/retry",
        "//server/util/status",
        "@com_github_jonboulle_clockwork//:clockwork",
        "@io_gorm_gorm//:gorm",
        "@org_golang_x_crypto//hkdf",
        "@org_golang_x_time//rate",
    ],
)

