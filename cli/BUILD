load("@rules_python//python:defs.bzl", "py_binary", "py_library")

# gazelle:exclude example_plugins/**/*.py
# gazelle:exclude plugins/**/*.py
# gazelle:default_visibility //cli:__subpackages__
package(default_visibility = ["//cli:__subpackages__"])

alias(
    name = "cli",
    actual = "//cli/cmd/bb",
)

genquery(
    name = "cli_enterprise_deps",
    expression = "filter(//enterprise, deps(//cli))",
    scope = ["//cli"],
    tags = ["manual"],
)

exports_files(
    srcs = ["install.sh"],
)

py_binary(
    name = "release",
    srcs = ["release.py"],
    visibility = ["//:__subpackages__"],
)

py_library(
    name = "cli_py_library",
    srcs = ["release.py"],
    visibility = ["//:__subpackages__"],
)
