load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day07_library",
    srcs = ["main.go", "day07.go"],
    importpath = "github.com/reusing-code/adventofcode/2022/day07",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day07",
    embed = [":go_day07_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day07_test",
    srcs = ["day07_test.go"],
    embed = [":go_day07_library"],
)
