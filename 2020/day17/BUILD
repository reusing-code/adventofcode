load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day17-lib",
    srcs = ["day17.cc"],
    hdrs = ["day17.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day17",
    srcs = ["main.cc"],
    deps = [
        ":day17-lib",
    ],
)

cc_test(
    name = "day17_test",
    size = "small",
    srcs = ["day17_test.cc"],
    deps = [
        ":day17-lib",
        "@googletest//:gtest_main",
    ],
)
