load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day9-lib",
    srcs = ["day9.cc"],
    hdrs = ["day9.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day9",
    srcs = ["main.cc"],
    deps = [
        ":day9-lib",
    ],
)

cc_test(
    name = "day9_test",
    size = "enormous",
    srcs = ["day9_test.cc"],
    deps = [
        ":day9-lib",
        "@googletest//:gtest_main",
    ],
)
