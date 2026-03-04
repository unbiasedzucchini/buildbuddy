GUEST_KERNEL_BUILDER_IMAGE = "docker://gcr.io/flame-public/guest-kernel-builder:latest"

def microvm_kernel(
        name,
        config,
        output,
        kernel_version,
        workdir,
        cpu_constraint):
    native.genrule(
        name = name,
        srcs = [config],
        outs = [output],
        cmd_bash = """
set -euo pipefail
mkdir -p .tmp
log_file="$$PWD/.tmp/%s.log"
if ! env \\
    TMPDIR="$$PWD/.tmp" \\
    KERNEL_VERSION=%s \\
    SKIP_UPLOAD=1 \\
    OUTPUT_FILE="$@" \\
    WORKDIR=%s \\
    $(location rebuild.sh) >"$$log_file" 2>&1; then
  cat "$$log_file" >&2
  exit 1
fi
""" % (name, kernel_version, workdir),
        exec_compatible_with = [
            cpu_constraint,
            "@platforms//os:linux",
        ],
        exec_properties = {
            "container-image": GUEST_KERNEL_BUILDER_IMAGE,
            "dockerNetwork": "bridge",
            "EstimatedComputeUnits": "16",
        },
        tags = ["manual"],
        target_compatible_with = [
            cpu_constraint,
            "@platforms//os:linux",
        ],
        tools = ["rebuild.sh"],
    )
