load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day18_library",
    srcs = ["main.go", "day18.go"],
    importpath = "github.com/reusing-code/adventofcode/2021/day18",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day18",
    embed = [":go_day18_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day18_test",
    srcs = ["day18_test.go"],
    embed = [":go_day18_library"],
)
