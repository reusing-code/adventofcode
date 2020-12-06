#include "day6.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day6, day6PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{{
                  "abc",  //
                  "",     //
                  "a",    //
                  "b",    //
                  "c",    //
                  "",     //
                  "ab",   //
                  "ac",   //
                  "",     //
                  "a",    //
                  "a",    //
                  "a",    //
                  "a",    //
                  "",     //
                  "b",    //
              },
              11};
  checkTC(tc, day6::puzzle1);
}

TEST(day6, day6PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{{
                  "abc",  //
                  "",     //
                  "a",    //
                  "b",    //
                  "c",    //
                  "",     //
                  "ab",   //
                  "ac",   //
                  "",     //
                  "a",    //
                  "a",    //
                  "a",    //
                  "a",    //
                  "",     //
                  "b",    //
              },
              6};
  checkTC(tc, day6::puzzle2);
}