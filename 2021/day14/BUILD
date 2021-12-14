load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day14_library",
    srcs = ["main.go", "day14.go"],
    importpath = "github.com/reusing-code/adventofcode/2021/day14",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day14",
    embed = [":go_day14_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day14_test",
    srcs = ["day14_test.go"],
    embed = [":go_day14_library"],
)
