load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

# Pre-generated from //proto:invocation_proto using protoc-gen-go + protoc-gen-go-vtproto.
# If you edit proto/invocation.proto, regenerate with:
#   bazel build //proto:invocation_go_proto
#   cp bazel-bin/proto/invocation_go_proto_/.../invocation/*.pb.go proto/invocation/
#   ./proto/invocation/update_hash.sh
go_library(
    name = "invocation",
    srcs = [
        "invocation.pb.go",
        "invocation_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/invocation",
    deps = [
        "//proto:acl_go_proto",
        "//proto:build_event_stream_go_proto",
        "//proto:cache_go_proto",
        "//proto:capability_go_proto",
        "//proto:command_line_go_proto",
        "//proto:context_go_proto",
        "//proto:invocation_status_go_proto",
        "//proto:stat_filter_go_proto",
        "//proto:target_go_proto",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@com_github_planetscale_vtprotobuf//types/known/durationpb",
        "@com_github_planetscale_vtprotobuf//types/known/timestamppb",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
    ],
)

go_test(
    name = "invocation_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:invocation.proto",
    ],
    rundir = ".",
)
