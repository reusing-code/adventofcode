#include "day11.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day11, day11PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "L.LL.LL.LL",  //
                  "LLLLLLL.LL",  //
                  "L.L.L..L..",  //
                  "LLLL.LL.LL",  //
                  "L.LL.LL.LL",  //
                  "L.LLLLL.LL",  //
                  "..L.L.....",  //
                  "LLLLLLLLLL",  //
                  "L.LLLLLL.L",  //
                  "L.LLLLL.LL",  //
              },
              37};
  checkTC(tc, day11::puzzle1);
}

TEST(day11, day11PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "L.LL.LL.LL",  //
                  "LLLLLLL.LL",  //
                  "L.L.L..L..",  //
                  "LLLL.LL.LL",  //
                  "L.LL.LL.LL",  //
                  "L.LLLLL.LL",  //
                  "..L.L.....",  //
                  "LLLLLLLLLL",  //
                  "L.LLLLLL.L",  //
                  "L.LLLLL.LL",  //
              },
              26};
  checkTC(tc, day11::puzzle2);
}