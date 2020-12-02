#include <fstream>
#include <iostream>
#include <string>

#include "day1.h"

int main(int argc, char const *argv[]) {
  if (argc < 2) {
    std::cerr << "Require input argument";
    return 0;
  }
  std::cout << "Running day1\n";;
  uint64_t number;
  std::vector<uint64_t> input;
  std::ifstream instr(argv[1]);
  while (instr >> number) {
    input.push_back(number);
  }

  auto result = day1::expense(input);
  std::cout << "result1: " << result << "\n";
  result = day1::expense3(input);
  std::cout << "result2: " << result << "\n";
  
  return 0;
}