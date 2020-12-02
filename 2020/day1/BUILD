load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day1-lib",
    srcs = ["day1.cc"],
    hdrs = ["day1.h"],
)

cc_binary(
    name = "day1",
    srcs = ["main.cc"],
    deps = [
        ":day1-lib",
    ],
)

cc_test(
    name = "day1_test",
    size = "enormous",
    srcs = ["day1_test.cc"],
    deps = [
        ":day1-lib",
        "@googletest//:gtest_main",
    ],
)
