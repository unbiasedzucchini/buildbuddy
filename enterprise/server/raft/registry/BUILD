load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "registry",
    srcs = ["registry.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/raft/registry",
    deps = [
        "//enterprise/server/raft/constants",
        "//proto:raft_go_proto",
        "//server/interfaces",
        "//server/util/kuberesolver",
        "//server/util/log",
        "//server/util/proto",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_hashicorp_serf//serf",
        "@com_github_lni_dragonboat_v4//config",
        "@com_github_lni_dragonboat_v4//raftio",
    ],
)

