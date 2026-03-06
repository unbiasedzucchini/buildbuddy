load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "platform",
    srcs = ["platform.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/platform",
    visibility = ["//visibility:public"],
    deps = [
        "//proto:firecracker_go_proto",
        "//proto:remote_execution_go_proto",
        "//proto:scheduler_go_proto",
        "//server/util/log",
        "//server/util/status",
        "@com_github_docker_go_units//:go-units",
        "@org_golang_google_grpc//metadata",
        "@org_golang_google_protobuf//encoding/protojson",
    ],
)

