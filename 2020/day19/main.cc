#include <fstream>
#include <iostream>
#include <string>

#include "day19.h"

int main(int argc, char const *argv[]) {
  if (argc < 2) {
    std::cerr << "Require input argument";
    return 0;
  }
  std::cout << "Running day19\n";
  std::string line;
  std::ifstream instr(argv[1]);
  std::vector<std::string> input;
  while (std::getline(instr, line)) {
    input.push_back(line);
  }
  int64_t result1 = day19::puzzle1(input);
  std::cout << "result1: " << result1 << "\n";
  int64_t result2 = day19::puzzle2(input);
  std::cout << "result2: " << result2 << "\n";

  return 0;
}