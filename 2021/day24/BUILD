load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day24_library",
    srcs = ["main.go", "day24.go"],
    importpath = "github.com/reusing-code/adventofcode/2021/day24",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day24",
    embed = [":go_day24_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day24_test",
    srcs = ["day24_test.go"],
    embed = [":go_day24_library"],
)
