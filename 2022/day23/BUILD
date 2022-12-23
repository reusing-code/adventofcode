load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_day23_library",
    srcs = ["main.go", "day23.go"],
    importpath = "github.com/reusing-code/adventofcode/2022/day23",
    deps = ["//gohelpers"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "day23",
    embed = [":go_day23_library"],
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_day23_test",
    srcs = ["day23_test.go"],
    embed = [":go_day23_library"],
)
