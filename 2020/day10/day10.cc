#include "day10.h"

#include <algorithm>
#include <array>

#include "helpers/helpers.h"

int64_t day10::puzzle1(const std::vector<std::string>& in) {
  auto intVec = helpers::parseIntVec(in);
  intVec.push_back(0);
  std::array<int64_t, 3> joltDiffs{0, 0, 1};
  std::sort(intVec.begin(), intVec.end());
  std::adjacent_find(intVec.begin(), intVec.end(),
                     [&joltDiffs](auto a, auto b) {
                       joltDiffs[b - a - 1]++;
                       return false;
                     });
  return joltDiffs[0] * joltDiffs[2];
}

int64_t day10::puzzle2(const std::vector<std::string>& in) { return 0; }