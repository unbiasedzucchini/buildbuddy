# Firecracker Guest Kernel Build + Upload

This directory contains the configs and scripts for building guest kernels used
by Firecracker.

Base configs are derived from Firecracker guest configs:
https://github.com/firecracker-microvm/firecracker/tree/main/resources/guest_configs

## Building and uploading the guest kernel

When updating guest kernel configs, you need to build them and manually
upload them to GCS.

There are two ways to do this, listed below. For both methods, it is
relatively safe to let the tools upload their outputs to GCS
automatically, since the GCS URLs contain a sha256 hash of the artifacts
that were built.

### Method 1 (old): Manual build + upload with `rebuild.sh`

`rebuild.sh` builds the guest kernel images and attempts to upload them to
GCS by default (requires `gsutil` to be installed and properly
authenticated).

When run from an x86_64 host, it will build and upload the x86_64 guest
kernel. When run from an arm64 host, it will build and upload the arm64
guest kernel.

```bash
cd enterprise/vmsupport/kernel
./rebuild.sh
```

To skip GCS upload, set `SKIP_UPLOAD=1`.

### Method 2 (recommended): multi-arch build and upload with bazel

Run the upload tool target (it builds and uploads all three kernels:
x86_64-v5.15, x86_64-v6.1, aarch64-v5.10):

```bash
# --config=remote enables RBE
# --config=target-linux-x86-exec-multiarch enables the BB arm64 platform in the build graph
# (not enabled by default at the time of writing)
bb run //enterprise/vmsupport/kernel/upload --config=remote --config=target-linux-x86-exec-multiarch
```

## Using the guest kernel

After building and uploading the guest kernel, update the
`org_kernel_git_linux_kernel-vmlinux*` targets in `deps.bzl` with the new
URLs and sha256 for the target(s) you intended to rebuild. This
information should get printed out by the build tools from the above
steps.
