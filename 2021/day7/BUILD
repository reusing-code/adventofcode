load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day7_library",
    srcs = ["main.go", "day7.go"],
    importpath = "github.com/reusing-code/adventofcode/2021/day7",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day7",
    embed = [":go_day7_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day7_test",
    srcs = ["day7_test.go"],
    embed = [":go_day7_library"],
)
