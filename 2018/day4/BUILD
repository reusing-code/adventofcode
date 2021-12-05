load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day4_library",
    srcs = ["main.go", "day4.go"],
    importpath = "github.com/reusing-code/adventofcode/2018/day4",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day4",
    embed = [":go_day4_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day4_test",
    srcs = ["day4_test.go"],
    embed = [":go_day4_library"],
)
