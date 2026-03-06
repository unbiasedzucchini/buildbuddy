load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "scorecard",
    srcs = ["scorecard.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/remote_cache/scorecard",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:build_event_stream_go_proto",
        "//proto:cache_go_proto",
        "//proto:pagination_go_proto",
        "//server/environment",
        "//server/util/paging",
        "//server/util/proto",
        "//server/util/status",
        "@org_golang_google_grpc//codes",
    ],
)

