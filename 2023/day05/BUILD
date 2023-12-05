load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day05_library",
    srcs = ["main.go", "day05.go"],
    importpath = "github.com/reusing-code/adventofcode/2023/day05",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day05",
    embed = [":go_day05_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day05_test",
    srcs = ["day05_test.go"],
    embed = [":go_day05_library"],
)
