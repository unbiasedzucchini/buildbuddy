load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

package(default_visibility = ["//visibility:public"])

go_library(
    name = "remote_execution",
    srcs = [
        "remote_execution.pb.go",
        "remote_execution_grpc.pb.go",
        "remote_execution_vtproto.pb.go",
    ],
    importpath = "github.com/buildbuddy-io/buildbuddy/proto/remote_execution",
    deps = [
        "//proto/scheduler",
        "//proto/semver",
        "@com_github_planetscale_vtprotobuf//protohelpers",
        "@com_github_planetscale_vtprotobuf//types/known/anypb",
        "@com_github_planetscale_vtprotobuf//types/known/durationpb",
        "@com_github_planetscale_vtprotobuf//types/known/timestamppb",
        "@com_github_planetscale_vtprotobuf//types/known/wrapperspb",
        "@com_google_cloud_go_longrunning//autogen/longrunningpb",
        "@org_golang_google_genproto_googleapis_api//annotations",
        "@org_golang_google_genproto_googleapis_rpc//status",
        "@org_golang_google_grpc//:grpc",
        "@org_golang_google_grpc//codes",
        "@org_golang_google_grpc//status",
        "@org_golang_google_protobuf//proto",
        "@org_golang_google_protobuf//reflect/protoreflect",
        "@org_golang_google_protobuf//runtime/protoimpl",
        "@org_golang_google_protobuf//types/known/anypb",
        "@org_golang_google_protobuf//types/known/durationpb",
        "@org_golang_google_protobuf//types/known/timestamppb",
        "@org_golang_google_protobuf//types/known/wrapperspb",
    ],
)

go_test(
    name = "remote_execution_drift_test",
    srcs = ["drift_test.go"],
    data = [
        ".proto_hash",
        "//proto:remote_execution.proto",
    ],
    rundir = ".",
)
