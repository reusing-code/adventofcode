#include "day19.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day19, day19PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "0: 4 1",        //
                  "1: 2 3 | 3 2",  //
                  "2: 4 4 | 5 5",  //
                  "3: 4 5 | 5 4",  //
                  "4: \"a\"",      //
                  "5: \"b\"",      //
                  "",              //
                  "ababb",         //
                  "babab",         //
                  "abbba",         //
                  "aaabb",         //
                  "aaaabb",        //
              },
              2};
  checkTC(tc, day19::puzzle1);
}

TEST(day19, day19PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{""}, 0};
  checkTC(tc, day19::puzzle2);
}