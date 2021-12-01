load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "day1.go",
        "main.go",
    ],
    importpath = "github.com/reusing-code/adventofcode/2021/day1",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day1",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["day1_test.go"],
    embed = [":go_default_library"],
)
