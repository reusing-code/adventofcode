load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day11_library",
    srcs = ["main.go", "day11.go"],
    importpath = "github.com/reusing-code/adventofcode/2022/day11",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day11",
    embed = [":go_day11_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day11_test",
    srcs = ["day11_test.go"],
    embed = [":go_day11_library"],
)
