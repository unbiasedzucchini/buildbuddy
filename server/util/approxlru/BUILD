load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "approxlru",
    srcs = ["approxlru.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/approxlru",
    visibility = ["//visibility:public"],
    deps = [
        "//server/util/log",
        "//server/util/status",
        "@com_github_jonboulle_clockwork//:clockwork",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_x_time//rate",
    ],
)

