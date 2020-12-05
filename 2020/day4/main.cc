#include <fstream>
#include <iostream>
#include <string>

#include "day4.h"

int main(int argc, char const *argv[]) {
  if (argc < 2) {
    std::cerr << "Require input argument";
    return 0;
  }
  std::cout << "Running day4\n";
  std::string line;
  std::ifstream instr(argv[1]);
  std::vector<std::string> input;
  while (std::getline(instr, line)) {
    input.push_back(line);
  }
  int result1 = day4::passports(input);
  std::cout << "result1: " << result1 << "\n";
  int result2 = day4::passports2(input);
  std::cout << "result2: " << result2 << "\n";

  return 0;
}