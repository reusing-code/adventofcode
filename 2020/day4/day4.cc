#include "day4.h"

#include <bitset>
#include <iostream>
#include <sstream>
#include <string>
#include <unordered_map>
#include <unordered_set>

const std::unordered_map<std::string, int> fieldMapping{{{"byr", 0},
                                                         {"iyr", 1},
                                                         {"eyr", 2},
                                                         {"hgt", 3},
                                                         {"hcl", 4},
                                                         {"ecl", 5},
                                                         {"pid", 6}}};

using FieldBitset = std::bitset<7>;

const std::unordered_set<std::string> allowedEyeColors{
    {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}};

int day4::passports(std::vector<std::string> in) {
  int result{0};

  FieldBitset bitset;
  for (const auto& str : in) {
    if (str == "") {
      if (bitset.all()) {
        result++;
      }
      bitset.reset();
      continue;
    }

    std::stringstream stream{str};
    while (stream.good()) {
      std::string val;
      stream >> val;
      auto sep = val.find(":");
      if (sep != std::string::npos) {
        auto key = val.substr(0, sep);
        auto idx = fieldMapping.find(key);
        if (idx != fieldMapping.end()) {
          bitset.set(idx->second);
        }
      }
    }
  }

  if (bitset.all()) {
    result++;
  }

  return result;
}

void checkValue(FieldBitset& bitset, const std::string& key,
                const std::string& value) {
  auto idx = fieldMapping.find(key);
  if (idx == fieldMapping.end()) {
    return;
  }
  auto checkYear = [](const std::string& str, int min, int max) -> bool {
    if (str.length() != 4) {
      return false;
    }
    int yr = std::stoi(str);
    if (yr >= min && yr <= max) {
      return true;
    }
    return false;
  };
  switch (idx->second) {
    case 0:  // byr
      if (checkYear(value, 1920, 2002)) {
        bitset.set(0);
      }
      break;
    case 1:  // iyr
      if (checkYear(value, 2010, 2020)) {
        bitset.set(1);
      }
      break;
    case 2:  // eyr
      if (checkYear(value, 2020, 2030)) {
        bitset.set(2);
      }
      break;
    case 3:  // hgt
    {
      int num;
      std::stringstream{value} >> num;
      if (value.substr(value.size() - 2) == "in") {
        if (num >= 59 && num <= 76) {
          bitset.set(3);
        }
      }
      if (value.substr(value.size() - 2) == "cm") {
        if (num >= 150 && num <= 193) {
          bitset.set(3);
        }
      }
    } break;
    case 4:  // hcl
    {
      auto itr = value.begin();
      if (*itr++ == '#') {
        for (; itr != value.end(); ++itr) {
          if ((*itr >= 'a' && *itr <= 'f') || (*itr >= '0' && *itr <= '9')) {
            bitset.set(4);
          }
        }
      }
    } break;
    case 5:  // ecl
      if (allowedEyeColors.find(value) != allowedEyeColors.end()) {
        bitset.set(5);
      }
      break;
    case 6:  // pid
      if (value.size() == 9) {
        for (char c : value) {
          if (c < '0' || c > '9') {
            break;
          }
        }
        bitset.set(6);
      }
      break;
  }
}

int day4::passports2(std::vector<std::string> in) {
  int result{0};

  FieldBitset bitset;
  for (const auto& str : in) {
    if (str == "") {
      if (bitset.all()) {
        result++;
      }
      bitset.reset();
      continue;
    }

    std::stringstream stream{str};
    while (stream.good()) {
      std::string val;
      stream >> val;
      auto sep = val.find(":");
      if (sep != std::string::npos) {
        auto key = val.substr(0, sep);
        auto value = val.substr(sep + 1);
        checkValue(bitset, key, value);
      }
    }
  }

  if (bitset.all()) {
    result++;
  }

  return result;
}
