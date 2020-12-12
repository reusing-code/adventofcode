#include "day8.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day8, day8PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{{
                  "nop +0",   //
                  "acc +1",   //
                  "jmp +4",   //
                  "acc +3",   //
                  "jmp -3",   //
                  "acc -99",  //
                  "acc +1",   //
                  "jmp -4",   //
                  "acc +6",   //
              },
              5};
  checkTC(tc, day8::puzzle1);
}

TEST(day8, day8PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{{
                  "nop +0",   //
                  "acc +1",   //
                  "jmp +4",   //
                  "acc +3",   //
                  "jmp -3",   //
                  "acc -99",  //
                  "acc +1",   //
                  "jmp -4",   //
                  "acc +6",   //
              },
              8};
  checkTC(tc, day8::puzzle2);
}