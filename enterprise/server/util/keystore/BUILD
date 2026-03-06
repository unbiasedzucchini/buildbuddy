load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "keystore",
    srcs = ["keystore.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/util/keystore",
    deps = [
        "//server/environment",
        "//server/util/status",
        "@org_golang_x_crypto//nacl/box",
    ],
)

