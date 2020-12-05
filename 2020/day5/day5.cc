#include "day5.h"

#include <algorithm>
#include <bitset>
#include <iostream>

int calcID(const std::string& str) {
  std::bitset<7> row;
  std::bitset<3> column;
  for (int i = 0; i < 7; ++i) {
    if (str[i] == 'B') {
      row.set(6 - i);
    }
  }
  for (int i = 0; i < 3; ++i) {
    if (str[7 + i] == 'R') {
      column.set(2 - i);
    }
  }
  return row.to_ulong() * 8 + column.to_ulong();
}

int day5::puzzle1(const std::vector<std::string>& in) {
  int max{0};
  for (const auto& str : in) {
    max = std::max(calcID(str), max);
  }
  return max;
}

int day5::puzzle2(const std::vector<std::string>& in) {
  std::vector<int> ids(in.size());
  std::generate(ids.begin(), ids.end(),
                [it = in.begin()]() mutable { return calcID(*it++); });
  std::sort(ids.begin(), ids.end());
  auto res = std::adjacent_find(ids.begin(), ids.end(),
                                [](int a, int b) { return b - a == 2; });
  return *res + 1;
}