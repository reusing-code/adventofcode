load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day6_library",
    srcs = ["main.go", "day6.go"],
    importpath = "github.com/reusing-code/adventofcode/2021/day6",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day6",
    embed = [":go_day6_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day6_test",
    srcs = ["day6_test.go"],
    embed = [":go_day6_library"],
)
