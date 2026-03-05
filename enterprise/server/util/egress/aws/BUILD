load("@io_bazel_rules_go//go:def.bzl", "go_binary")

package(default_visibility = ["//enterprise:__subpackages__"])

genrule(
    name = "raw-ranges",
    outs = ["raw-ranges.json"],
    cmd_bash = """
    curl --fail --silent --show-error --location \
      https://ip-ranges.amazonaws.com/ip-ranges.json > "$@"
    """,
    local = True,
)

go_binary(
    name = "ranges",
    srcs = ["ranges.go"],
    visibility = ["//enterprise/server/util/egress:__pkg__"],
)
