load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_library(
    name = "grpc",
    srcs = glob(["*.go"]),
    importpath = "github.com/TheGrizzlyDev/rules_integration_tests/core/go/grpc",
    deps = [
        ":containers_grpc_lib",
        "@org_golang_google_grpc//:grpc",
    ],
    visibility = [
        "//core/go:__pkg__",
    ],
)

go_proto_library(
    name = "containers_grpc_lib",
    importpath = "github.com/TheGrizzlyDev/rules_integration_tests/proto/containers",
    protos = [
        "//core/proto:containers",
    ],
    compilers = [
        "@io_bazel_rules_go//proto:go_grpc",
    ]
)