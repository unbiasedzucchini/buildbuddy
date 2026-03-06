load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "hostedrunner",
    srcs = ["hostedrunner.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/hostedrunner",
    deps = [
        "//enterprise/server/experiments",
        "//enterprise/server/remote_execution/operation",
        "//enterprise/server/remote_execution/snaputil",
        "//enterprise/server/util/ci_runner_env",
        "//enterprise/server/util/ci_runner_util",
        "//enterprise/server/workflow/config",
        "//proto:invocation_status_go_proto",
        "//proto:remote_execution_go_proto",
        "//proto:runner_go_proto",
        "//server/endpoint_urls/build_buddy_url",
        "//server/endpoint_urls/cache_api_url",
        "//server/endpoint_urls/events_api_url",
        "//server/endpoint_urls/remote_exec_api_url",
        "//server/environment",
        "//server/remote_cache/cachetools",
        "//server/remote_cache/digest",
        "//server/tables",
        "//server/util/bazel_request",
        "//server/util/db",
        "//server/util/git",
        "//server/util/log",
        "//server/util/platform",
        "//server/util/prefix",
        "//server/util/rexec",
        "//server/util/status",
        "@com_github_google_uuid//:uuid",
        "@com_google_cloud_go_longrunning//autogen/longrunningpb",
        "@in_gopkg_yaml_v2//:yaml_v2",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//types/known/durationpb",
    ],
)

