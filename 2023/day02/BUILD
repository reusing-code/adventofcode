load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day02_library",
    srcs = ["main.go", "day02.go"],
    importpath = "github.com/reusing-code/adventofcode/2023/day02",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day02",
    embed = [":go_day02_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day02_test",
    srcs = ["day02_test.go"],
    embed = [":go_day02_library"],
)
