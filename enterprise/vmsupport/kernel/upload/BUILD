load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_binary(
    name = "upload",
    data = [
        "//enterprise/vmsupport/kernel:guest_kernel_arm64",
        "//enterprise/vmsupport/kernel:guest_kernel_x86_64_v5_15",
        "//enterprise/vmsupport/kernel:guest_kernel_x86_64_v6_1",
    ],
    embed = [":upload_lib"],
    tags = ["manual"],
    visibility = ["//visibility:public"],
    x_defs = {
        "arm64V5_10KernelRlocationpath": "$(rlocationpath //enterprise/vmsupport/kernel:guest_kernel_arm64)",
        "x86V5_15KernelRlocationpath": "$(rlocationpath //enterprise/vmsupport/kernel:guest_kernel_x86_64_v5_15)",
        "x86V6_1KernelRlocationpath": "$(rlocationpath //enterprise/vmsupport/kernel:guest_kernel_x86_64_v6_1)",
    },
)

go_library(
    name = "upload_lib",
    srcs = ["upload.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/vmsupport/kernel/upload",
    tags = ["manual"],
    deps = [
        "//server/util/flag",
        "//server/util/log",
        "@io_bazel_rules_go//go/runfiles",
    ],
)
