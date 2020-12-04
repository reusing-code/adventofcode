#include "day3.h"

int day3::trees(std::vector<std::string> in) {
  int result{0};
  std::vector<std::vector<bool>> map;
  for (const auto& str : in) {
    std::vector<bool> row;
    for (char c : str) {
      row.push_back(c == '#');
    }
    map.push_back(row);
  }
  int r{0};
  int c{0};
  while (r < map.size()) {
    if (map[r][c]) {
      result++;
    }
    c = (c + 3) % map[r].size();
    r++;
  }
  return result;
}

int countTrees(const std::vector<std::vector<bool>>& map, int rows, int cols) {
  int result{0};
  int r{0};
  int c{0};
  while (r < map.size()) {
    if (map[r][c]) {
      result++;
    }
    c = (c + cols) % map[r].size();
    r += rows;
  }
  return result;
}

int day3::trees2(std::vector<std::string> in) {
  int result{1};
  std::vector<std::vector<bool>> map;
  for (const auto& str : in) {
    std::vector<bool> row;
    for (char c : str) {
      row.push_back(c == '#');
    }
    map.push_back(row);
  }
  result *= countTrees(map, 1, 1);
  result *= countTrees(map, 1, 3);
  result *= countTrees(map, 1, 5);
  result *= countTrees(map, 1, 7);
  result *= countTrees(map, 2, 1);

  return result;
}
