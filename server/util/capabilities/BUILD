load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "capabilities",
    srcs = ["capabilities.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/capabilities",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:capability_go_proto",
        "//server/interfaces",
        "//server/util/authutil",
        "//server/util/status",
    ],
)

