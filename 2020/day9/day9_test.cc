#include "day9.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day9, day9PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{
      {
          "35",  "20",  "15",  "25",  "47",  "40",  "62",  "55",  "65",  "95",
          "102", "117", "150", "182", "127", "219", "299", "277", "309", "576",
      },
      127};
  checkTC(tc, std::bind(day9::puzzle1, std::placeholders::_1, 5));
}

TEST(day9, day9PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{
      {
          "35",  "20",  "15",  "25",  "47",  "40",  "62",  "55",  "65",  "95",
          "102", "117", "150", "182", "127", "219", "299", "277", "309", "576",
      },
      62};
  checkTC(tc, std::bind(day9::puzzle2, std::placeholders::_1, 5));
}