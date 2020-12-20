#include "day14.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day14, day14PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",  //
                  "mem[8] = 11",                                  //
                  "mem[7] = 101",                                 //
                  "mem[8] = 0",                                   //
              },
              165};
  checkTC(tc, day14::puzzle1);
}

TEST(day14, day14PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "mask = 000000000000000000000000000000X1001X",  //
                  "mem[42] = 100",                                //
                  "mask = 00000000000000000000000000000000X0XX",  //
                  "mem[26] = 1",                                  //
              },
              208};
  checkTC(tc, day14::puzzle2);
}