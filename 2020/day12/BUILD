load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day12-lib",
    srcs = ["day12.cc"],
    hdrs = ["day12.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day12",
    srcs = ["main.cc"],
    deps = [
        ":day12-lib",
    ],
)

cc_test(
    name = "day12_test",
    size = "small",
    srcs = ["day12_test.cc"],
    deps = [
        ":day12-lib",
        "@googletest//:gtest_main",
    ],
)
