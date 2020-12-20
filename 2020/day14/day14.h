#pragma once

#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>

class day14 {
 public:
  static int64_t puzzle1(const std::vector<std::string>& in);
  static int64_t puzzle2(const std::vector<std::string>& in);

  struct Program {
    int64_t maskOne{0};
    int64_t maskZero{0};
    std::vector<int64_t> flippyBits;

    std::unordered_map<int64_t, int64_t> mem;

    static Program parse(const std::vector<std::string>& in);
    static Program parse2(const std::vector<std::string>& in);
  };

 private:
};