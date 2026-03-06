load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "auditlog",
    srcs = ["auditlog.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/auditlog",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:auditlog_go_proto",
        "//proto:capability_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/real_environment",
        "//server/util/capabilities",
        "//server/util/claims",
        "//server/util/clickhouse/schema",
        "//server/util/clientip",
        "//server/util/db",
        "//server/util/log",
        "//server/util/proto",
        "//server/util/query_builder",
        "//server/util/random",
        "//server/util/status",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

