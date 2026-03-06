load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "schema",
    srcs = ["schema.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/clickhouse/schema",
    visibility = ["//visibility:public"],
    deps = [
        "//server/build_event_protocol/invocation_format",
        "//server/tables",
        "//server/usage/sku",
        "//server/util/log",
        "//server/util/status",
        "@io_gorm_gorm//:gorm",
    ],
)

