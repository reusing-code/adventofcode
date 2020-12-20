#include "day14.h"

#include <bits/stdint-uintn.h>

#include <algorithm>
#include <bitset>
#include <numeric>

#include "helpers/helpers.h"

int64_t day14::puzzle1(const std::vector<std::string>& in) {
  Program prog = Program::parse(in);
  return std::accumulate(prog.mem.begin(), prog.mem.end(), 0ll,
                         [](int64_t a, const auto& b) { return a + b.second; });
}

int64_t day14::puzzle2(const std::vector<std::string>& in) {
  Program prog = Program::parse2(in);
  return std::accumulate(prog.mem.begin(), prog.mem.end(), 0ll,
                         [](int64_t a, const auto& b) { return a + b.second; });
}

day14::Program day14::Program::parse(const std::vector<std::string>& in) {
  Program prog;
  for (const auto& str : in) {
    if (str.starts_with("mask")) {
      const auto& mask = str.substr(7);
      std::bitset<36> ones;
      std::bitset<36> zeroes;
      for (size_t i = 0; i < mask.size(); i++) {
        if (mask[mask.size() - i - 1] == '1') {
          ones.set(i);
        } else if (mask[mask.size() - i - 1] == '0') {
          zeroes.set(i);
        }
      }
      prog.maskOne = ones.to_ullong();
      prog.maskZero = zeroes.flip().to_ullong();
    } else if (str.starts_with("mem")) {
      int64_t adr = std::stoll(str.substr(4));
      int64_t val = std::stoll(str.substr(str.find("=") + 1));
      prog.mem[adr] = (val | prog.maskOne) & prog.maskZero;
    }
  }

  return prog;
}

bool increase(std::vector<bool>& bs) {
  for (std::size_t i = 0; i != bs.size(); ++i) {
    bs[i] = !bs[i];
    if (bs[i] == true) {
      return true;
    }
  }
  return false;  // overflow
}

template <typename T>
std::vector<std::vector<T>> PowerSet(const std::vector<T>& v) {
  std::vector<std::vector<T>> result;
  std::vector<bool> bitset(v.size());

  do {
    std::vector<T> cur;
    for (std::size_t i = 0; i != v.size(); ++i) {
      if (bitset[i]) {
        cur.push_back(v[i]);
      }
    }
    result.push_back(cur);
  } while (increase(bitset));
  return result;
}

day14::Program day14::Program::parse2(const std::vector<std::string>& in) {
  Program prog;
  for (const auto& str : in) {
    if (str.starts_with("mask")) {
      const auto& mask = str.substr(7);
      std::bitset<36> ones;
      std::bitset<36> zeroes;
      prog.flippyBits.clear();
      for (size_t i = 0; i < mask.size(); i++) {
        if (mask[mask.size() - i - 1] == '1') {
          ones.set(i);
        } else if (mask[mask.size() - i - 1] == '0') {
          zeroes.set(i);
        } else {
          prog.flippyBits.push_back(i);
        }
      }
      prog.maskOne = ones.to_ullong();
      prog.maskZero = zeroes.flip().to_ullong();
    } else if (str.starts_with("mem")) {
      std::bitset<36> baseAdr{static_cast<uint64_t>(std::stoll(str.substr(4))) |
                              prog.maskOne};
      int64_t val = std::stoll(str.substr(str.find("=") + 1));
      auto power = PowerSet(prog.flippyBits);
      for (const auto& var : power) {
        std::bitset<36> adr = baseAdr;
        for (auto it : prog.flippyBits) {
          adr.reset(it);
        }
        for (auto it : var) {
          adr.set(it);
        }
        prog.mem[adr.to_ullong()] = val;
      }
    }
  }

  return prog;
}