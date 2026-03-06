load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "ociruntime",
    srcs = ["ociruntime.go"],
    data = [":crun"],
    embedsrcs = [
        "hosts",
        "seccomp.json",
        "tini",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/remote_execution/containers/ociruntime",
    target_compatible_with = ["@platforms//os:linux"],
    visibility = ["//visibility:public"],
    x_defs = {
        "crunRlocationpath": "$(rlocationpath :crun)",
    },
    deps = [
        "//enterprise/server/remote_execution/block_io",
        "//enterprise/server/remote_execution/cgroup",
        "//enterprise/server/remote_execution/commandutil",
        "//enterprise/server/remote_execution/container",
        "//enterprise/server/remote_execution/executor_auth",
        "//enterprise/server/util/oci",
        "//enterprise/server/util/ociconv",
        "//proto:remote_execution_go_proto",
        "//proto:scheduler_go_proto",
        "//server/environment",
        "//server/interfaces",
        "//server/util/claims",
        "//server/util/disk",
        "//server/util/flag",
        "//server/util/fsync",
        "//server/util/hash",
        "//server/util/log",
        "//server/util/networking",
        "//server/util/platform",
        "//server/util/proto",
        "//server/util/status",
        "//server/util/statusz",
        "//server/util/unixcred",
        "//server/util/uuid",
        "//third_party/singleflight",
        "@com_github_google_go_containerregistry//pkg/v1:pkg",
        "@com_github_opencontainers_runtime_spec//specs-go",
        "@io_bazel_rules_go//go/runfiles",
        "@org_golang_x_sync//errgroup",
        "@org_golang_x_sys//unix",
    ],
)

genrule(
    name = "tini_binary",
    srcs = ["//enterprise/server/cmd/executor:tini"],
    outs = ["tini"],
    cmd = "cp $(SRCS) $@",
)

alias(
    name = "crun",
    actual = select({
        "@platforms//cpu:x86_64": "@com_github_containers_crun_crun-linux-amd64//file:crun",
        "@platforms//cpu:aarch64": "@com_github_containers_crun_crun-linux-arm64//file:crun",
    }),
)

alias(
    name = "busybox",
    actual = select({
        "@platforms//cpu:x86_64": "@net_busybox_busybox-linux-amd64//file:busybox",
        "@platforms//cpu:aarch64": "@net_busybox_busybox-linux-arm64//file:busybox",
    }),
)
