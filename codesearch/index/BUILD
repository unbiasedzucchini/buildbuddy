load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "index",
    srcs = [
        "index.go",
        "merger.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/codesearch/index",
    visibility = ["//visibility:public"],
    deps = [
        "//codesearch/performance",
        "//codesearch/posting",
        "//codesearch/types",
        "//server/util/log",
        "//server/util/status",
        "@com_github_cockroachdb_pebble//:pebble",
        "@com_github_xiam_s_expr//ast",
        "@com_github_xiam_s_expr//parser",
        "@org_golang_x_sync//errgroup",
    ],
)

