load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "docker",
    srcs = ["docker.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/containers/docker",
    deps = [
        "//enterprise/server/remote_execution/commandutil",
        "//enterprise/server/remote_execution/container",
        "//enterprise/server/remote_execution/executorplatform",
        "//enterprise/server/util/oci",
        "//proto:remote_execution_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/platform",
        "//server/util/random",
        "//server/util/status",
        "@com_github_docker_docker//api/types/container",
        "@com_github_docker_docker//api/types/image",
        "@com_github_docker_docker//api/types/registry",
        "@com_github_docker_docker//client",
        "@com_github_docker_docker//pkg/stdcopy",
        "@com_github_docker_go_units//:go-units",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_x_sync//errgroup",
    ],
)

