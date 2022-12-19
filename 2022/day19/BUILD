load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day19_library",
    srcs = ["main.go", "day19.go"],
    importpath = "github.com/reusing-code/adventofcode/2022/day19",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day19",
    embed = [":go_day19_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day19_test",
    srcs = ["day19_test.go"],
    embed = [":go_day19_library"],
)
