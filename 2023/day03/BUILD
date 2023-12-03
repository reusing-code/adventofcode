load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day03_library",
    srcs = ["main.go", "day03.go"],
    importpath = "github.com/reusing-code/adventofcode/2023/day03",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day03",
    embed = [":go_day03_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day03_test",
    srcs = ["day03_test.go"],
    embed = [":go_day03_library"],
)
