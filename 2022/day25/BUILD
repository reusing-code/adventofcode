load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day25_library",
    srcs = ["main.go", "day25.go"],
    importpath = "github.com/reusing-code/adventofcode/2022/day25",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day25",
    embed = [":go_day25_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day25_test",
    srcs = ["day25_test.go"],
    embed = [":go_day25_library"],
)
