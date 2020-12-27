#include "day17.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day17, day17PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  ".#.",  //
                  "..#",  //
                  "###",  //
              },
              112};
  checkTC(tc, day17::puzzle1);
}

TEST(day17, day17PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  ".#.",  //
                  "..#",  //
                  "###",  //
              },
              848};
  checkTC(tc, day17::puzzle2);
}