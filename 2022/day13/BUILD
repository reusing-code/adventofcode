load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day13_library",
    srcs = ["main.go", "day13.go"],
    importpath = "github.com/reusing-code/adventofcode/2022/day13",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day13",
    embed = [":go_day13_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day13_test",
    srcs = ["day13_test.go"],
    embed = [":go_day13_library"],
)
