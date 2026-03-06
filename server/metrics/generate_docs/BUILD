load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "generate_docs_lib",
    srcs = ["generate_docs.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/server/metrics/generate_docs",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "generate_docs",
    embed = [":generate_docs_lib"],
    visibility = ["//:__subpackages__"],
)

genrule(
    name = "generate_mdx",
    srcs = ["//server/metrics:metrics.go"],
    outs = ["docs.mdx"],
    cmd_bash = """
    BAZEL_BINDIR=$(BINDIR) $(location :generate_docs) \
        --prettier_path=$(location //tools/prettier) \
        --input_path=$(SRCS) \
        --output_path=$@
    """,
    tools = [
        ":generate_docs",
        "//tools/prettier",
    ],
    visibility = [
        "//website:__pkg__",
    ],
)
