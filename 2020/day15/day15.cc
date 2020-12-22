#include "day15.h"

#include <algorithm>

#include "helpers/helpers.h"

using SeenMap = std::unordered_map<int64_t, std::vector<int64_t>>;

void add(SeenMap &map, int64_t number, uint64_t turn)

{
  map[number].push_back(static_cast<int64_t>(turn));
}

int64_t newNum(SeenMap &map, int64_t lastNum)

{
  int64_t num{0};

  auto find = map.find(lastNum);

  if (find != map.end() && find->second.size() > 1) {
    num = find->second.back() - find->second.at(find->second.size() - 2);
  }

  return num;
}

int64_t day15::puzzle1(const std::vector<std::string> &in) {
  SeenMap lastSeen;

  std::vector<int64_t> startValues = helpers::splitInt(in.front(), ',');

  uint64_t turn = 1;

  for (; turn <= startValues.size(); ++turn) {
    add(lastSeen, startValues[turn - 1], turn);
  }

  int64_t lastValue = startValues.back();

  for (; turn <= 2020; ++turn) {
    lastValue = newNum(lastSeen, lastValue);

    add(lastSeen, lastValue, turn);
  }

  return lastValue;
}

int64_t day15::puzzle2(const std::vector<std::string> &in) {
  SeenMap lastSeen;

  std::vector<int64_t> startValues = helpers::splitInt(in.front(), ',');

  uint64_t turn = 1;

  for (; turn <= startValues.size(); ++turn) {
    add(lastSeen, startValues[turn - 1], turn);
  }

  int64_t lastValue = startValues.back();

  for (; turn <= 30000000; ++turn) {
    lastValue = newNum(lastSeen, lastValue);

    add(lastSeen, lastValue, turn);
  }

  return lastValue;
}