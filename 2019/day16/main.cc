#include <fstream>
#include <iostream>
#include <string>

#include "day16.h"

int main(int argc, char const *argv[]) {
  if (argc < 2) {
    std::cerr << "Require input argument";
    return 0;
  }
  std::string input;
  std::ifstream instr(argv[1]);
  std::getline(instr, input);

  FFT fft({0, 1, 0, -1});
  auto result = fft.run(input, 100);
  std::cout << "Result: " << result << std::endl;
  result = fft.runPartTwo(input, 100);
  std::cout << "Result2: " << result << std::endl;
  return 0;
}