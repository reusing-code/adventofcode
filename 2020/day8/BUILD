load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day8-lib",
    srcs = ["day8.cc"],
    hdrs = ["day8.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day8",
    srcs = ["main.cc"],
    deps = [
        ":day8-lib",
    ],
)

cc_test(
    name = "day8_test",
    size = "enormous",
    srcs = ["day8_test.cc"],
    deps = [
        ":day8-lib",
        "@googletest//:gtest_main",
    ],
)
