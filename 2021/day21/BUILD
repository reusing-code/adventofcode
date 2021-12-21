load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day21_library",
    srcs = ["main.go", "day21.go"],
    importpath = "github.com/reusing-code/adventofcode/2021/day21",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day21",
    embed = [":go_day21_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day21_test",
    srcs = ["day21_test.go"],
    embed = [":go_day21_library"],
)
