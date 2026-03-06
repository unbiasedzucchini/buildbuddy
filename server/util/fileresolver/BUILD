load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fileresolver",
    srcs = ["fileresolver.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/fileresolver",
    visibility = ["//visibility:public"],
    x_defs = {
        "fileresolverGoRlocation": "$(rlocationpath fileresolver.go)",
    },
    deps = ["@io_bazel_rules_go//go/runfiles"],
)

