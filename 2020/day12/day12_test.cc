#include "day12.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day12, day12PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "F10",  //
                  "N3",   //
                  "F7",   //
                  "R90",  //
                  "F11",  //
              },
              25};
  checkTC(tc, day12::puzzle1);
}

TEST(day12, day12PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "F10",  //
                  "N3",   //
                  "F7",   //
                  "R90",  //
                  "F11",  //
              },
              286};
  checkTC(tc, day12::puzzle2);
}