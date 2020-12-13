#include "day10.h"

#include <algorithm>
#include <array>
#include <memory>

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

int64_t day10::puzzle2(const std::vector<std::string>& in) {
  if (in.empty()) {
    return 0;
  }
  auto intVec = helpers::parseIntVec(in);
  std::sort(intVec.begin(), intVec.end());
  auto lookup = std::make_unique<std::vector<int64_t>>(intVec.back() + 1, 0);
  (*lookup)[0] = 1;
  for (auto it : intVec) {
    int64_t combinations{(*lookup)[it - 1]};
    if (it > 1) {
      combinations += (*lookup)[it - 2];
    }
    if (it > 2) {
      combinations += (*lookup)[it - 3];
    }
    (*lookup)[it] = combinations;
  }
  return lookup->back();
}