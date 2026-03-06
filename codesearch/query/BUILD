load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "query",
    srcs = [
        "regexp.go",
        "regexp_query.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/codesearch/query",
    visibility = ["//visibility:public"],
    deps = [
        "//codesearch/dfa",
        "//codesearch/filters",
        "//codesearch/token",
        "//codesearch/types",
        "//server/util/log",
        "//server/util/status",
    ],
)

