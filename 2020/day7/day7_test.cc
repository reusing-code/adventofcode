#include "day7.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

auto checkTC = [](const auto& tc, auto checkfun) {
  auto& in = std::get<0>(tc);
  auto& want = std::get<1>(tc);
  auto got = checkfun(in);
  EXPECT_EQ(want, got);
};

TEST(day7, day7PartOne) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{
      {
          "light red bags contain 1 bright white bag, 2 muted yellow bags.",  //
          "dark orange bags contain 3 bright white bags, 4 muted yellow bags.",  //
          "bright white bags contain 1 shiny gold bag.",                      //
          "muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",  //
          "shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",   //
          "dark olive bags contain 3 faded blue bags, 4 dotted black bags.",  //
          "vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",  //
          "faded blue bags contain no other bags.",    //
          "dotted black bags contain no other bags.",  //
      },
      4};
  checkTC(tc, day7::puzzle1);
}

TEST(day7, day7PartTwo) {
  using TestCase = std::tuple<std::vector<std::string>, int>;
  TestCase tc{{
                  "shiny gold bags contain 2 dark red bags.",      //
                  "dark red bags contain 2 dark orange bags.",     //
                  "dark orange bags contain 2 dark yellow bags.",  //
                  "dark yellow bags contain 2 dark green bags.",   //
                  "dark green bags contain 2 dark blue bags.",     //
                  "dark blue bags contain 2 dark violet bags.",    //
                  "dark violet bags contain no other bags.",       //
              },
              126};
  checkTC(tc, day7::puzzle2);
}