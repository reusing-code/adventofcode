load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day17_library",
    srcs = ["main.go", "day17.go"],
    importpath = "github.com/reusing-code/adventofcode/2022/day17",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day17",
    embed = [":go_day17_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day17_test",
    srcs = ["day17_test.go"],
    embed = [":go_day17_library"],
)
