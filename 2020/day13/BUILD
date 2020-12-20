load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day13-lib",
    srcs = ["day13.cc"],
    hdrs = ["day13.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day13",
    srcs = ["main.cc"],
    deps = [
        ":day13-lib",
    ],
)

cc_test(
    name = "day13_test",
    size = "small",
    srcs = ["day13_test.cc"],
    deps = [
        ":day13-lib",
        "@googletest//:gtest_main",
    ],
)
