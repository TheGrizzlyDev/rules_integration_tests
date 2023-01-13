load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_binary(
    name = "go",
    srcs = glob(["*.go"]),
    deps = [
        ":containers_grpc_lib",
    ],
)

go_proto_library(
    name = "containers_grpc_lib",
    importpath = "github.com/TheGrizzlyDev/rules_integration_tests/proto/containers",
    protos = [
        "//core/proto:containers",
    ],
)