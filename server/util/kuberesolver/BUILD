load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "kuberesolver",
    srcs = ["kuberesolver.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/util/kuberesolver",
    visibility = ["//visibility:public"],
    deps = [
        "//server/util/log",
        "//server/util/status",
        "@io_k8s_api//core/v1:core",
        "@io_k8s_apimachinery//pkg/api/errors",
        "@io_k8s_apimachinery//pkg/apis/meta/v1:meta",
        "@io_k8s_apimachinery//pkg/watch",
        "@io_k8s_client_go//kubernetes",
        "@io_k8s_client_go//rest",
        "@org_golang_google_grpc//resolver",
    ],
)

