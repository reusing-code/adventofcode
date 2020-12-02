#include <fstream>
#include <iostream>
#include <string>

#include "@DAY@.h"

int main(int argc, char const *argv[]) {
  if (argc < 2) {
    std::cerr << "Require input argument";
    return 0;
  }
  std::cout << "Running @DAY@\n";
  std::string input;
  std::ifstream instr(argv[1]);
  std::getline(instr, input);

  return 0;
}