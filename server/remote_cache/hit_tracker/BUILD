load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "hit_tracker",
    srcs = ["hit_tracker.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/remote_cache/hit_tracker",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:cache_go_proto",
        "//proto:remote_execution_go_proto",
        "//proto:resource_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/tables",
        "//server/util/alert",
        "//server/util/log",
        "//server/util/proto",
        "//server/util/rexec",
        "//server/util/status",
        "//server/util/usageutil",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_google_genproto_googleapis_rpc//status",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

