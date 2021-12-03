load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day3_library",
    srcs = ["main.go", "day3.go"],
    importpath = "github.com/reusing-code/adventofcode/2018/day3",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day3",
    embed = [":go_day3_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day3_test",
    srcs = ["day3_test.go"],
    embed = [":go_day3_library"],
)
