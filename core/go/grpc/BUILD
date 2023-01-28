load("@rules_go//go:def.bzl", "go_library")
load("@rules_go//proto:def.bzl", "go_proto_library")

go_library(
    name = "grpc",
    srcs = glob(["*.go"]),
    importpath = "github.com/TheGrizzlyDev/rules_integration_tests/core/go/grpc",
    deps = [
        ":containers_grpc_lib",
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
        "@rules_go//proto:go_grpc",
    ]
)