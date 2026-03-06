load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "ociregistry",
    srcs = ["ociregistry.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/ociregistry",
    visibility = ["//visibility:public"],
    deps = [
        "//enterprise/server/util/ocicache",
        "//proto:ociregistry_go_proto",
        "//proto:remote_execution_go_proto",
        "//server/environment",
        "//server/http/httpclient",
        "//server/real_environment",
        "//server/util/log",
        "//server/util/prefix",
        "//server/util/status",
        "@com_github_google_go_containerregistry//pkg/name",
        "@com_github_google_go_containerregistry//pkg/v1:pkg",
        "@org_golang_google_genproto_googleapis_bytestream//:bytestream",
    ],
)

