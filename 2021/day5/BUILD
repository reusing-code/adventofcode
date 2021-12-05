load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day5_library",
    srcs = ["main.go", "day5.go"],
    importpath = "github.com/reusing-code/adventofcode/2021/day5",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day5",
    embed = [":go_day5_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day5_test",
    srcs = ["day5_test.go"],
    embed = [":go_day5_library"],
)
