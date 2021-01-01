load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day18-lib",
    srcs = ["day18.cc"],
    hdrs = ["day18.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day18",
    srcs = ["main.cc"],
    deps = [
        ":day18-lib",
    ],
)

cc_test(
    name = "day18_test",
    size = "small",
    srcs = ["day18_test.cc"],
    deps = [
        ":day18-lib",
        "@googletest//:gtest_main",
    ],
)
