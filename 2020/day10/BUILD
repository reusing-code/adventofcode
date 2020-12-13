load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day10-lib",
    srcs = ["day10.cc"],
    hdrs = ["day10.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day10",
    srcs = ["main.cc"],
    deps = [
        ":day10-lib",
    ],
)

cc_test(
    name = "day10_test",
    size = "small",
    srcs = ["day10_test.cc"],
    deps = [
        ":day10-lib",
        "@googletest//:gtest_main",
    ],
)
