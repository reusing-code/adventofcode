load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day2-lib",
    srcs = ["day2.cc"],
    hdrs = ["day2.h"],
)

cc_binary(
    name = "day2",
    srcs = ["main.cc"],
    deps = [
        ":day2-lib",
    ],
)

cc_test(
    name = "day2_test",
    size = "enormous",
    srcs = ["day2_test.cc"],
    deps = [
        ":day2-lib",
        "@googletest//:gtest_main",
    ],
)
