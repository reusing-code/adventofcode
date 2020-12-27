#include "day17.h"

#include <algorithm>
#include <queue>

#include "helpers/helpers.h"

int64_t day17::puzzle1(const std::vector<std::string>& in) {
  std::unordered_set<Coordinate> active;
  int x = 0;
  for (const auto& str : in) {
    int y = 0;
    for (char c : str) {
      if (c == '#') {
        active.emplace(x, y, 0);
      }
      ++y;
    }
    ++x;
  }
  for (int i = 0; i < 6; ++i) {
    active = doStep(active);
  }
  return active.size();
}

int64_t day17::puzzle2(const std::vector<std::string>& in) {
  std::unordered_set<Coordinate> active;
  int x = 0;
  for (const auto& str : in) {
    int y = 0;
    for (char c : str) {
      if (c == '#') {
        active.emplace(x, y, 0);
      }
      ++y;
    }
    ++x;
  }
  for (int i = 0; i < 6; ++i) {
    active = doStep2(active);
  }
  return active.size();
}

std::unordered_set<day17::Coordinate> day17::doStep(
    const std::unordered_set<Coordinate>& old) {
  std::unordered_set<Coordinate> result;
  std::unordered_set<Coordinate> checked;
  std::queue<Coordinate> queue;

  for (const auto& it : old) {
    queue.push(it);
    checked.insert(it);
  }

  while (!queue.empty()) {
    const auto coord = queue.front();
    checked.insert(coord);
    bool active = old.find(coord) != old.end();
    int count{0};
    for (int64_t dx = -1; dx <= 1; ++dx) {
      for (int64_t dy = -1; dy <= 1; ++dy) {
        for (int64_t dz = -1; dz <= 1; ++dz) {
          if (dx == 0 && dy == 0 && dz == 0) {
            continue;
          }
          Coordinate newCoord{coord.x + dx, coord.y + dy, coord.z + dz};
          if (old.find(newCoord) != old.end()) {
            ++count;
          }
          if (active && checked.find(newCoord) == checked.end()) {
            checked.insert(newCoord);
            queue.push(newCoord);
          }
        }
      }
    }
    if (active) {
      if (count == 2 || count == 3) {
        result.insert(coord);
      }
    } else {
      if (count == 3) {
        result.insert(coord);
      }
    }

    queue.pop();
  }

  return result;
}

std::unordered_set<day17::Coordinate> day17::doStep2(
    const std::unordered_set<Coordinate>& old) {
  std::unordered_set<Coordinate> result;
  std::unordered_set<Coordinate> checked;
  std::queue<Coordinate> queue;

  for (const auto& it : old) {
    queue.push(it);
    checked.insert(it);
  }

  while (!queue.empty()) {
    const auto coord = queue.front();
    checked.insert(coord);
    bool active = old.find(coord) != old.end();
    int count{0};
    for (int64_t dx = -1; dx <= 1; ++dx) {
      for (int64_t dy = -1; dy <= 1; ++dy) {
        for (int64_t dz = -1; dz <= 1; ++dz) {
          for (int64_t dw = -1; dw <= 1; ++dw) {
            if (dx == 0 && dy == 0 && dz == 0 && dw == 0) {
              continue;
            }
            Coordinate newCoord{coord.x + dx, coord.y + dy, coord.z + dz,
                                coord.w + dw};
            if (old.find(newCoord) != old.end()) {
              ++count;
            }
            if (active && checked.find(newCoord) == checked.end()) {
              checked.insert(newCoord);
              queue.push(newCoord);
            }
          }
        }
      }
    }
    if (active) {
      if (count == 2 || count == 3) {
        result.insert(coord);
      }
    } else {
      if (count == 3) {
        result.insert(coord);
      }
    }

    queue.pop();
  }

  return result;
}

namespace std {
std::size_t hash<day17::Coordinate>::operator()(
    day17::Coordinate const& c) const noexcept {
  std::size_t h1 = std::hash<int64_t>{}(c.x);
  std::size_t h2 = std::hash<int64_t>{}(c.y);
  std::size_t h3 = std::hash<int64_t>{}(c.z);
  std::size_t h4 = std::hash<int64_t>{}(c.w);
  return h1 ^ (h2 << 1) ^ (h3 << 2) ^ (h4 << 3);
}
}  // namespace std

bool operator==(const day17::Coordinate& lhs, const day17::Coordinate& rhs) {
  return lhs.x == rhs.x && lhs.y == rhs.y && lhs.z == rhs.z && lhs.w == rhs.w;
}