load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day15-lib",
    srcs = ["day15.cc"],
    hdrs = ["day15.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day15",
    srcs = ["main.cc"],
    deps = [
        ":day15-lib",
    ],
)

cc_test(
    name = "day15_test",
    size = "enormous",
    srcs = ["day15_test.cc"],
    deps = [
        ":day15-lib",
        "@googletest//:gtest_main",
    ],
)
