load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "experiments",
    srcs = ["experiments.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/experiments",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:group_go_proto",
        "//server/interfaces",
        "//server/real_environment",
        "//server/util/bazel_request",
        "//server/util/claims",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/proto",
        "//server/util/region",
        "//server/util/statusz",
        "@com_github_open_feature_go_sdk//openfeature",
        "@com_github_open_feature_go_sdk_contrib_providers_flagd//pkg",
        "@org_golang_google_protobuf//encoding/protojson",
    ],
)

