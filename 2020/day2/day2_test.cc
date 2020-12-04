#include "day2.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

TEST(day2, day2PartOne) {
  using TestTable = std::vector<std::tuple<std::vector<std::string>, int>>;
  TestTable tt{{{"1-3 a: abcde",  //
                 "1-3 b: cdefg",  //
                 "2-9 c: ccccccccc"},
                2}};

  for (auto& tc : tt) {
    auto& in = std::get<0>(tc);
    auto& want = std::get<1>(tc);
    auto got = day2::password(in);
    EXPECT_EQ(want, got);
  }
}

TEST(day2, day2PartTwo) {
  using TestTable = std::vector<std::tuple<std::vector<std::string>, int>>;
  TestTable tt{{{"1-3 a: abcde",  //
                 "1-3 b: cdefg",  //
                 "2-9 c: ccccccccc"},
                1}};

  for (auto& tc : tt) {
    auto& in = std::get<0>(tc);
    auto& want = std::get<1>(tc);
    auto got = day2::password2(in);
    EXPECT_EQ(want, got);
  }
}