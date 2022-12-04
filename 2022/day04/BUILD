load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day04_library",
    srcs = ["main.go", "day04.go"],
    importpath = "github.com/reusing-code/adventofcode/2022/day04",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day04",
    embed = [":go_day04_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day04_test",
    srcs = ["day04_test.go"],
    embed = [":go_day04_library"],
)
