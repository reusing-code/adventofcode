#include "day5.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day5, day5PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}, 820};
  checkTC(tc, day5::puzzle1);
}

TEST(day5, day5PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{{""}, 0};
  checkTC(tc, day5::puzzle2);
}