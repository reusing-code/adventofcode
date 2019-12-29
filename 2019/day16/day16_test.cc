#include "day16.h"

#include <gtest/gtest.h>

#include <tuple>
#include <vector>

TEST(FFT, FFTPartOne) {
  using TestTable = std::vector<std::tuple<std::string, int, std::string>>;
  TestTable tt{
      {"12345678", 1, "48226158"},
      {"12345678", 2, "34040438"},
      {"12345678", 3, "03415518"},
      {"12345678", 4, "01029498"},
      {"80871224585914546619083218645595", 100, "24176176"},
      {"19617804207202209144916044189917", 100, "73745418"},
      {"69317163492948606335995924319873", 100, "52432133"},
  };

  FFT fft({0, 1, 0, -1});

  for (auto& tc : tt) {
    auto& in = std::get<0>(tc);
    auto& iterations = std::get<1>(tc);
    auto& want = std::get<2>(tc);
    auto got = fft.run(in, iterations);
    EXPECT_EQ(want, got);
  }
}

TEST(FFT, FFTPartTwo) {
  using TestTable = std::vector<std::tuple<std::string, int, std::string>>;
  TestTable tt{
      {"03036732577212944063491565474664", 100, "84462026"},
      {"02935109699940807407585447034323", 100, "78725270"},
      {"03081770884921959731165446850517", 100, "53553731"},
  };

  FFT fft({0, 1, 0, -1});

  for (auto& tc : tt) {
    auto& in = std::get<0>(tc);
    auto& iterations = std::get<1>(tc);
    auto& want = std::get<2>(tc);
    auto got = fft.runPartTwo(in, iterations);
    EXPECT_EQ(want, got);
  }
}