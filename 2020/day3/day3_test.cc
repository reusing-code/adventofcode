#include "day3.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

TEST(day3, day3PartOne) {
  using TestTable = std::vector<std::tuple<std::vector<std::string>, int>>;
  TestTable tt{{{
                    "..##.......",  //
                    "#...#...#..",  //
                    ".#....#..#.",  //
                    "..#.#...#.#",  //
                    ".#...##..#.",  //
                    "..#.##.....",  //
                    ".#.#.#....#",  //
                    ".#........#",  //
                    "#.##...#...",  //
                    "#...##....#",  //
                    ".#..#...#.#"   //
                },
                7}};

  for (auto& tc : tt) {
    auto& in = std::get<0>(tc);
    auto& want = std::get<1>(tc);
    auto got = day3::trees(in);
    EXPECT_EQ(want, got);
  }
}

TEST(day3, day3PartTwo) {
  using TestTable = std::vector<std::tuple<std::vector<std::string>, int>>;
  TestTable tt{{{
                    "..##.......",  //
                    "#...#...#..",  //
                    ".#....#..#.",  //
                    "..#.#...#.#",  //
                    ".#...##..#.",  //
                    "..#.##.....",  //
                    ".#.#.#....#",  //
                    ".#........#",  //
                    "#.##...#...",  //
                    "#...##....#",  //
                    ".#..#...#.#"   //
                },
                336}};

  for (auto& tc : tt) {
    auto& in = std::get<0>(tc);
    auto& want = std::get<1>(tc);
    auto got = day3::trees2(in);
    EXPECT_EQ(want, got);
  }
}