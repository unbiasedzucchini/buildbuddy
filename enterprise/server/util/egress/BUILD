load("@io_bazel_rules_go//go:def.bzl", "go_binary")

package(default_visibility = ["//enterprise:__subpackages__"])

genrule(
    name = "aws_ranges_csv",
    srcs = ["//enterprise/server/util/egress/aws:raw-ranges"],
    outs = ["aws_ranges.csv"],
    cmd_bash = """
    ./$(location //enterprise/server/util/egress/aws:ranges) \
      --input "$(location //enterprise/server/util/egress/aws:raw-ranges)" \
      --output "$@"
    """,
    local = True,
    tools = ["//enterprise/server/util/egress/aws:ranges"],
)

genrule(
    name = "gcp_ranges_csv",
    srcs = ["//enterprise/server/util/egress/gcp:raw-ranges"],
    outs = ["gcp_ranges.csv"],
    cmd_bash = """
    ./$(location //enterprise/server/util/egress/gcp:ranges) \
      --input "$(location //enterprise/server/util/egress/gcp:raw-ranges)" \
      --output "$@"
    """,
    local = True,
    tools = ["//enterprise/server/util/egress/gcp:ranges"],
)

go_binary(
    name = "egress",
    srcs = ["egress.go"],
    embedsrcs = [
        ":aws_ranges_csv",
        ":gcp_ranges_csv",
    ],
)
