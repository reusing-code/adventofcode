load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day15_library",
    srcs = ["main.go", "day15.go"],
    importpath = "github.com/reusing-code/adventofcode/2023/day15",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day15",
    embed = [":go_day15_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day15_test",
    srcs = ["day15_test.go"],
    embed = [":go_day15_library"],
)
