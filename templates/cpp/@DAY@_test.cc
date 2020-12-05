#include "@DAY@.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC =
    [](const auto& tc, auto checkfun) {
      auto& in = std::get<0>(tc);
      auto& want = std::get<1>(tc);
      auto got = checkfun(in);
      EXPECT_EQ(want, got);
    }

TEST(@DAY@, @DAY@PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{{""}, 0};
  checkTC(tc, @DAY@::puzzle1);
}

TEST(@DAY@, @DAY@PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{{""}, 0};
  checkTC(tc, @DAY@::puzzle2);
}