load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day6-lib",
    srcs = ["day6.cc"],
    hdrs = ["day6.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day6",
    srcs = ["main.cc"],
    deps = [
        ":day6-lib",
    ],
)

cc_test(
    name = "day6_test",
    size = "small",
    srcs = ["day6_test.cc"],
    deps = [
        ":day6-lib",
        "@googletest//:gtest_main",
    ],
)
