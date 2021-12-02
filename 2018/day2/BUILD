load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day2_library",
    srcs = ["main.go", "day2.go"],
    importpath = "github.com/reusing-code/adventofcode/2018/day2",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day2",
    embed = [":go_day2_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day2_test",
    srcs = ["day2_test.go"],
    embed = [":go_day2_library"],
)
