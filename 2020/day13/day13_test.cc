#include "day13.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day13, day13PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{"939", "7,13,x,x,59,x,31,19"}, 295};
  checkTC(tc, day13::puzzle1);
}

TEST(day13, day13PartTwo) {
  using TestCases = std::vector<std::tuple<std::vector<std::string>, int64_t>>;
  TestCases tcs{
      {{"939", "7,13,x,x,59,x,31,19"}, 1068781},
      {{"939", "17,x,13,19"}, 3417},
      {{"939", "67,7,59,61"}, 754018},
      {{"939", "67,x,7,59,61"}, 779210},
      {{"939", "67,7,x,59,61"}, 1261476},
      {{"939", "1789,37,47,1889"}, 1202161486},
  };
  for (const auto& tc : tcs) {
    checkTC(tc, day13::puzzle2);
  }
}