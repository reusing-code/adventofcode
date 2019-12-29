#include "day16.h"

#include <algorithm>
#include <iostream>
#include <sstream>
#include <string>

FFT::FFT(const std::vector<char>& basePattern) : basePattern_(basePattern) {}

std::string FFT::run(const std::string& input, uint32_t iterations) {
  offset_ = 0;
  currentValue_.clear();
  for (uint32_t i = 0; i < input.length(); ++i) {
    currentValue_.push_back(std::atoi(input.substr(i, 1).c_str()));
  }

  for (uint32_t i = 0; i < iterations; ++i) {
    iteration();
  }

  std::stringstream ss;
  for (auto& val : currentValue_) {
    ss << static_cast<int>(val);
  }

  return ss.str().substr(0, 8);
}

void FFT::iteration() {
  std::vector<char> newValue(currentValue_.size() - offset_);
  for (uint32_t i = offset_; i < currentValue_.size(); ++i) {
    int sum{0};
    for (uint32_t j = i; j < currentValue_.size(); ++j) {
      sum += getSummand(i, j);
    }
    newValue[i - offset_] = (std::abs(sum % 10));
  }
  std::copy(newValue.begin(), newValue.end(), currentValue_.begin() + offset_);
}

void FFT::iteration2() {
  int sum{0};
  for (uint32_t i = currentValue_.size() - 1; i >= offset_; --i) {
    sum += currentValue_[i];
    currentValue_[i] = (std::abs(sum % 10));
  }
}

std::string FFT::runPartTwo(const std::string& input, uint32_t iterations) {
  currentValue_.clear();
  for (uint32_t j = 0; j < 10000; ++j) {
    for (uint32_t i = 0; i < input.length(); ++i) {
      currentValue_.push_back(std::atoi(input.substr(i, 1).c_str()));
    }
  }

  offset_ = 0;
  for (size_t i = 0; i < 7; ++i) {
    offset_ = offset_ * 10 + currentValue_[i];
  }

  for (uint32_t i = 0; i < iterations; ++i) {
    iteration2();
  }

  std::stringstream ss;
  for (size_t i = offset_; i < offset_ + 8; ++i) {
    ss << static_cast<int>(currentValue_[i]);
  }

  return ss.str();
}

int FFT::getSummand(int row, int column) {
  int factor = basePattern_[((column + 1) / (row + 1)) % basePattern_.size()];
  return factor * currentValue_[column];
}
