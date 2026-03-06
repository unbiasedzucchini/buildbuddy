load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

exports_files(srcs = ["vmexec.go"])

go_library(
    name = "vmexec",
    srcs = ["vmexec.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/vmexec",
    target_compatible_with = [
        "@platforms//os:linux",
    ],
    deps = [
        "//enterprise/server/remote_execution/commandutil",
        "//enterprise/server/util/vsock",
        "//proto:remote_execution_go_proto",
        "//proto:vmexec_go_proto",
        "//server/remote_cache/digest",
        "//server/util/fsutil",
        "//server/util/healthcheck",
        "//server/util/log",
        "//server/util/networking",
        "//server/util/retry",
        "//server/util/status",
        "@com_github_elastic_gosigar//:gosigar",
        "@com_github_miekg_dns//:dns",
        "@com_github_tklauser_go_sysconf//:go-sysconf",
        "@com_github_vishvananda_netlink//:netlink",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//health/grpc_health_v1",
        "@org_golang_google_grpc//status",
        "@org_golang_x_sys//unix",
    ],
)

