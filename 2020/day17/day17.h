#pragma once

#include <iostream>
#include <string>
#include <unordered_set>
#include <vector>

class day17 {
 public:
  static int64_t puzzle1(const std::vector<std::string>& in);
  static int64_t puzzle2(const std::vector<std::string>& in);

  struct Coordinate {
    int64_t x{0};
    int64_t y{0};
    int64_t z{0};
    int64_t w{0};

    Coordinate(int64_t x_, int64_t y_, int64_t z_, int64_t w_ = 0)
        : x(x_), y(y_), z(z_), w(w_) {}
  };

 private:
  static std::unordered_set<Coordinate> doStep(
      const std::unordered_set<Coordinate>& old);
  static std::unordered_set<Coordinate> doStep2(
      const std::unordered_set<Coordinate>& old);
};

namespace std {
template <>
struct hash<day17::Coordinate> {
  std::size_t operator()(day17::Coordinate const& c) const noexcept;
};
}  // namespace std

bool operator==(const day17::Coordinate& lhs, const day17::Coordinate& rhs);