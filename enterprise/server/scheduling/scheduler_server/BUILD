load("@io_bazel_rules_go//go:def.bzl", "go_library")

# gazelle:default_visibility //enterprise:__subpackages__
package(default_visibility = [
    "//enterprise:__subpackages__",
])

go_library(
    name = "scheduler_server",
    srcs = ["scheduler_server.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/scheduling/scheduler_server",
    deps = [
        "//enterprise/server/experiments",
        "//enterprise/server/remote_execution/action_merger",
        "//enterprise/server/tasksize",
        "//enterprise/server/util/ci_runner_util",
        "//proto:capability_go_proto",
        "//proto:remote_execution_go_proto",
        "//proto:scheduler_go_proto",
        "//proto:trace_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/real_environment",
        "//server/remote_execution/config",
        "//server/resources",
        "//server/scheduling/scheduler_server/config",
        "//server/util/alert",
        "//server/util/authutil",
        "//server/util/background",
        "//server/util/bazel_request",
        "//server/util/error_util",
        "//server/util/grpc_client",
        "//server/util/log",
        "//server/util/perms",
        "//server/util/platform",
        "//server/util/proto",
        "//server/util/random",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_go_redis_redis_v8//:redis",
        "@com_github_jonboulle_clockwork//:clockwork",
        "@com_github_prometheus_client_golang//prometheus",
        "@org_golang_google_grpc//peer",
        "@org_golang_google_protobuf//encoding/prototext",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

