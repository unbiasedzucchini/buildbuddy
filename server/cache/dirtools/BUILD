load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "dirtools",
    srcs = ["dirtools.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/cache/dirtools",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:remote_execution_go_proto",
        "//proto:resource_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/remote_cache/cachetools",
        "//server/remote_cache/digest",
        "//server/util/alert",
        "//server/util/claims",
        "//server/util/fastcopy",
        "//server/util/fspath",
        "//server/util/ioutil",
        "//server/util/log",
        "//server/util/proto",
        "//server/util/rpcutil",
        "//server/util/status",
        "//third_party/singleflight",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_google_genproto_googleapis_bytestream//:bytestream",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_x_sync//errgroup",
        "@org_golang_x_sync//semaphore",
    ],
)

