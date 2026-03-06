load("@io_bazel_rules_go//go:def.bzl", "go_library")

package(default_visibility = ["//enterprise:__subpackages__"])

go_library(
    name = "oci",
    srcs = ["oci.go"],
    importpath = "github.com/buildbuddy-io/buildbuddy/enterprise/server/util/oci",
    deps = [
        "//enterprise/server/oci/ocifetcher",
        "//enterprise/server/util/ocicache",
        "//enterprise/server/util/ocimanifest",
        "//proto:oci_fetcher_go_proto",
        "//proto:registry_go_proto",
        "//proto:remote_execution_go_proto",
        "//server/environment",
        "//server/http/httpclient",
        "//server/util/authutil",
        "//server/util/claims",
        "//server/util/flag",
        "//server/util/log",
        "//server/util/lru",
        "//server/util/platform",
        "//server/util/status",
        "//server/util/tracing",
        "@com_github_distribution_reference//:reference",
        "@com_github_google_go_containerregistry//pkg/authn",
        "@com_github_google_go_containerregistry//pkg/name",
        "@com_github_google_go_containerregistry//pkg/v1:pkg",
        "@com_github_google_go_containerregistry//pkg/v1/partial",
        "@com_github_google_go_containerregistry//pkg/v1/remote",
        "@com_github_google_go_containerregistry//pkg/v1/remote/transport",
        "@com_github_google_go_containerregistry//pkg/v1/types",
        "@com_github_jonboulle_clockwork//:clockwork",
        "@org_golang_google_genproto_googleapis_bytestream//:bytestream",
    ],
)

