load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day16-lib",
    srcs = ["day16.cc"],
    hdrs = ["day16.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day16",
    srcs = ["main.cc"],
    deps = [
        ":day16-lib",
    ],
)

cc_test(
    name = "day16_test",
    size = "small",
    srcs = ["day16_test.cc"],
    deps = [
        ":day16-lib",
        "@googletest//:gtest_main",
    ],
)
