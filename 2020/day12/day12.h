#pragma once

#include <iostream>
#include <string>
#include <vector>

class day12 {
 public:
  static int64_t puzzle1(const std::vector<std::string>& in);
  static int64_t puzzle2(const std::vector<std::string>& in);

  struct Nav {
    int64_t x{0};
    int64_t y{0};
    int dir{1};

    void doStep(const std::string& instr);
  };

  struct Nav2 {
    struct Pos {
      int64_t x{0};
      int64_t y{0};
    };
    Pos ship;
    Pos waypoint{10, 1};

    void doStep(const std::string& instr);
    void rotate(int dir);
  };

 private:
};