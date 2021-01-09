load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library", "cc_test")

cc_library(
    name = "day19-lib",
    srcs = ["day19.cc"],
    hdrs = ["day19.h"],
    deps = [
        "//helpers",
    ],
)

cc_binary(
    name = "day19",
    srcs = ["main.cc"],
    deps = [
        ":day19-lib",
    ],
)

cc_test(
    name = "day19_test",
    size = "small",
    srcs = ["day19_test.cc"],
    deps = [
        ":day19-lib",
        "@googletest//:gtest_main",
    ],
)
