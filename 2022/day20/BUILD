load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day20_library",
    srcs = ["main.go", "day20.go"],
    importpath = "github.com/reusing-code/adventofcode/2022/day20",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day20",
    embed = [":go_day20_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day20_test",
    srcs = ["day20_test.go"],
    embed = [":go_day20_library"],
)
