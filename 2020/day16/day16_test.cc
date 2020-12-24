#include "day16.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day16, day16PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{
                  "class: 1-3 or 5-7",     //
                  "row: 6-11 or 33-44",    //
                  "seat: 13-40 or 45-50",  //
                  "",                      //
                  "your ticket:",          //
                  "7,1,14",                //
                  "",                      //
                  "nearby tickets:",       //
                  "7,3,47",                //
                  "40,4,50",               //
                  "55,2,20",               //
                  "38,6,12",               //
              },
              71};
  checkTC(tc, day16::puzzle1);
}

TEST(day16, day16PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int64_t>;
  TestCase tc{{""}, 0};
  checkTC(tc, day16::puzzle2);
}