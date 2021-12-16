load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day16_library",
    srcs = ["main.go", "day16.go"],
    importpath = "github.com/reusing-code/adventofcode/2021/day16",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day16",
    embed = [":go_day16_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day16_test",
    srcs = ["day16_test.go"],
    embed = [":go_day16_library"],
)
