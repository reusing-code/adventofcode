load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day06_library",
    srcs = ["main.go", "day06.go"],
    importpath = "github.com/reusing-code/adventofcode/2023/day06",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day06",
    embed = [":go_day06_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day06_test",
    srcs = ["day06_test.go"],
    embed = [":go_day06_library"],
)
