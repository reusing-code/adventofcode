#include "day1.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

TEST(day1, day1PartOne) {
  using TestTable = std::vector<std::tuple<std::vector<uint64_t>, uint64_t>>;
  TestTable tt{{{1721, 979, 366, 299, 675, 1456}, 514579}};

  for (auto& tc : tt) {
    auto& in = std::get<0>(tc);
    auto& want = std::get<1>(tc);
    auto got = day1::expense(in);
    EXPECT_EQ(want, got);
  }
}

TEST(day1, day1PartTwo) {
  using TestTable = std::vector<std::tuple<std::vector<uint64_t>, uint64_t>>;
  TestTable tt{{{1721, 979, 366, 299, 675, 1456}, 241861950}};

  for (auto& tc : tt) {
    auto& in = std::get<0>(tc);
    auto& want = std::get<1>(tc);
    auto got = day1::expense3(in);
    EXPECT_EQ(want, got);
  }
}