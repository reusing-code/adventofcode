load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day3-lib",
    srcs = ["day3.cc"],
    hdrs = ["day3.h"],
)

cc_binary(
    name = "day3",
    srcs = ["main.cc"],
    deps = [
        ":day3-lib",
    ],
)

cc_test(
    name = "day3_test",
    size = "small",
    srcs = ["day3_test.cc"],
    deps = [
        ":day3-lib",
        "@googletest//:gtest_main",
    ],
)
