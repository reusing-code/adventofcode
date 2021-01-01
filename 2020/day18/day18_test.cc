#include "day18.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day18, day18PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "2 * 3 + (4 * 5)",                                  //
                  "5 + (8 * 3 + 9 + 3 * 4 * 3)",                      //
                  "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",        //
                  "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",  //
              },
              26 + 437 + 12240 + 13632};
  checkTC(tc, day18::puzzle1);
}

TEST(day18, day18PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "2 * 3 + (4 * 5)",                                  //
                  "5 + (8 * 3 + 9 + 3 * 4 * 3)",                      //
                  "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",        //
                  "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",  //
              },
              46 + 1445 + 669060 + 23340};
  checkTC(tc, day18::puzzle2);
}