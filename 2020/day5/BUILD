load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day5-lib",
    srcs = ["day5.cc"],
    hdrs = ["day5.h"],
)

cc_binary(
    name = "day5",
    srcs = ["main.cc"],
    deps = [
        ":day5-lib",
    ],
)

cc_test(
    name = "day5_test",
    size = "small",
    srcs = ["day5_test.cc"],
    deps = [
        ":day5-lib",
        "@googletest//:gtest_main",
    ],
)
