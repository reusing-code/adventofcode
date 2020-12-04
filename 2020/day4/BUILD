load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day4-lib",
    srcs = ["day4.cc"],
    hdrs = ["day4.h"],
)

cc_binary(
    name = "day4",
    srcs = ["main.cc"],
    deps = [
        ":day4-lib",
    ],
)

cc_test(
    name = "day4_test",
    size = "enormous",
    srcs = ["day4_test.cc"],
    deps = [
        ":day4-lib",
        "@googletest//:gtest_main",
    ],
)
