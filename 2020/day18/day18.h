#pragma once

#include <iostream>
#include <string>
#include <vector>

class day18 {
 public:
  static int64_t puzzle1(const std::vector<std::string>& in);
  static int64_t puzzle2(const std::vector<std::string>& in);

 private:
  enum class Op { NONE, ADD, MUL };

  static int64_t solveEquation(const std::string& in);
  static int64_t solveEquation2(const std::string& in);
};