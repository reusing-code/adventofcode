load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day11-lib",
    srcs = ["day11.cc"],
    hdrs = ["day11.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day11",
    srcs = ["main.cc"],
    deps = [
        ":day11-lib",
    ],
)

cc_test(
    name = "day11_test",
    size = "small",
    srcs = ["day11_test.cc"],
    deps = [
        ":day11-lib",
        "@googletest//:gtest_main",
    ],
)
