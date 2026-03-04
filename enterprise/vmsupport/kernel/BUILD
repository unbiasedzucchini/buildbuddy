load("//server/util/bb:platform_transition.bzl", "linux_arm64_alias", "linux_x86_64_alias")
load(":defs.bzl", "microvm_kernel")

package(default_visibility = ["//enterprise:__subpackages__"])

microvm_kernel(
    name = "rebuild_guest_kernel_x86_64_v5_15",
    config = "microvm-kernel-x86_64-v5.15.config",
    cpu_constraint = "@platforms//cpu:x86_64",
    kernel_version = "v5.15",
    output = "vmlinux-x86_64-v5.15",
    workdir = "/tmp/buildbuddy-guest-kernel-x86_64-v5.15",
)

microvm_kernel(
    name = "rebuild_guest_kernel_x86_64_v6_1",
    config = "microvm-kernel-x86_64-v6.1.config",
    cpu_constraint = "@platforms//cpu:x86_64",
    kernel_version = "v6.1",
    output = "vmlinux-x86_64-v6.1",
    workdir = "/tmp/buildbuddy-guest-kernel-x86_64-v6.1",
)

microvm_kernel(
    name = "rebuild_guest_kernel_aarch64_v5_10",
    config = "microvm-kernel-aarch64-v5.10.config",
    cpu_constraint = "@platforms//cpu:arm64",
    kernel_version = "v5.10",
    output = "vmlinux-aarch64-v5.10",
    workdir = "/tmp/buildbuddy-guest-kernel-aarch64-v5.10",
)

linux_x86_64_alias(
    name = "guest_kernel_x86_64_v5_15",
    actual = ":rebuild_guest_kernel_x86_64_v5_15",
    tags = ["manual"],
)

linux_arm64_alias(
    name = "guest_kernel_arm64",
    actual = ":rebuild_guest_kernel_aarch64_v5_10",
    tags = ["manual"],
)

# EXPERIMENTAL: v6.1 guest kernel (x86_64 only). Not used in prod yet.
linux_x86_64_alias(
    name = "guest_kernel_x86_64_v6_1",
    actual = ":rebuild_guest_kernel_x86_64_v6_1",
    tags = ["manual"],
)
