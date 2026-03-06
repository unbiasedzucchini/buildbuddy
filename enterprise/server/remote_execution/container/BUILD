load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "container",
    srcs = ["container.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/container",
    deps = [
        "//enterprise/server/remote_execution/block_io",
        "//enterprise/server/remote_execution/executor_auth",
        "//enterprise/server/remote_execution/operation",
        "//enterprise/server/util/oci",
        "//proto:firecracker_go_proto",
        "//proto:remote_execution_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/metrics",
        "//server/util/alert",
        "//server/util/authutil",
        "//server/util/background",
        "//server/util/flag",
        "//server/util/hash",
        "//server/util/log",
        "//server/util/platform",
        "//server/util/status",
        "//server/util/tracing",
        "//server/util/unixcred",
        "@com_github_jonboulle_clockwork//:clockwork",
        "@com_github_prometheus_client_golang//prometheus",
        "@io_opentelemetry_go_otel//attribute",
        "@io_opentelemetry_go_otel_trace//:trace",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

