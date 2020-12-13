load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day7-lib",
    srcs = ["day7.cc"],
    hdrs = ["day7.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day7",
    srcs = ["main.cc"],
    deps = [
        ":day7-lib",
    ],
)

cc_test(
    name = "day7_test",
    size = "small",
    srcs = ["day7_test.cc"],
    deps = [
        ":day7-lib",
        "@googletest//:gtest_main",
    ],
)
