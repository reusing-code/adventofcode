load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day8_library",
    srcs = ["main.go", "day8.go"],
    importpath = "github.com/reusing-code/adventofcode/2021/day8",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day8",
    embed = [":go_day8_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day8_test",
    srcs = ["day8_test.go"],
    embed = [":go_day8_library"],
)
