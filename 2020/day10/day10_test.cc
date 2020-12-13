#include "day10.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day10, day10PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{
      {
          "28", "33", "18", "42", "31", "14", "46", "20", "48", "47", "24",
          "23", "49", "45", "19", "38", "39", "11", "1",  "32", "25", "35",
          "8",  "17", "7",  "9",  "4",  "2",  "34", "10", "3",
      },
      220};
  checkTC(tc, day10::puzzle1);
}

TEST(day10, day10PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{""}, 0};
  checkTC(tc, day10::puzzle2);
}