load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day14-lib",
    srcs = ["day14.cc"],
    hdrs = ["day14.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day14",
    srcs = ["main.cc"],
    deps = [
        ":day14-lib",
    ],
)

cc_test(
    name = "day14_test",
    size = "small",
    srcs = ["day14_test.cc"],
    deps = [
        ":day14-lib",
        "@googletest//:gtest_main",
    ],
)
