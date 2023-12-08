load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day08_library",
    srcs = ["main.go", "day08.go"],
    importpath = "github.com/reusing-code/adventofcode/2023/day08",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day08",
    embed = [":go_day08_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day08_test",
    srcs = ["day08_test.go"],
    embed = [":go_day08_library"],
)
