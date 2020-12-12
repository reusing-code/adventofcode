#include "day9.h"

#include <numeric>
#include <set>

#include "helpers/helpers.h"

bool checkSum(const std::set<int64_t>& preamble, int64_t value) {
  auto front = preamble.begin();
  auto back = --preamble.end();
  while (front != back) {
    int64_t sum = *front + *back;
    if (sum == value) {
      return true;
    }
    if (sum < value) {
      ++front;
    } else {
      --back;
    }
  }
  return false;
}

int64_t day9::puzzle1(const std::vector<std::string>& in, int preambleSize) {
  auto intVec = helpers::parseIntVec(in);

  std::set<int64_t> preamble{intVec.begin(), intVec.begin() + preambleSize};

  for (auto it = intVec.begin() + preambleSize; it != intVec.end(); ++it) {
    if (!checkSum(preamble, *it)) {
      return *it;
    }

    preamble.erase(*(it - preambleSize));
    preamble.insert(*it);
  }

  return -1;
}

int64_t day9::puzzle2(const std::vector<std::string>& in, int preambleSize) {
  int64_t targetValue = puzzle1(in, preambleSize);
  auto intVec = helpers::parseIntVec(in);

  auto front = intVec.begin();
  auto back = ++intVec.begin();
  while (true) {
    int64_t sum = std::accumulate(front, back, 0ll);
    if (sum == targetValue) {
      auto minmax = std::minmax_element(front, back);
      return *minmax.first + *minmax.second;
    }
    if (sum < targetValue) {
      ++back;
    } else {
      ++front;
    }
  }
  return 0;
}