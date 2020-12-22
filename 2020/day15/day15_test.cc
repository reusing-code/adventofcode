#include "day15.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day15, day15PartOne) {
  using TestCases = std::vector<std::tuple<std::vector<std::string>, int64_t>>;
  TestCases tcs{
      {{"1,3,2"}, 1},  {{"2,1,3"}, 10},  {{"1,2,3"}, 27},
      {{"2,3,1"}, 78}, {{"3,2,1"}, 438}, {{"3,1,2"}, 1836},
  };
  for (const auto& tc : tcs) {
    checkTC(tc, day15::puzzle1);
  }
}

TEST(day15, day15PartTwo) {
  using TestCases = std::vector<std::tuple<std::vector<std::string>, int64_t>>;
  TestCases tcs{
      {{"1,3,2"}, 2578},    {{"2,1,3"}, 3544142}, {{"1,2,3"}, 261214},
      {{"2,3,1"}, 6895259}, {{"3,2,1"}, 18},      {{"3,1,2"}, 362},
  };
  for (const auto& tc : tcs) {
    checkTC(tc, day15::puzzle2);
  }
}