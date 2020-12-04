#include "day2.h"

#include <iostream>
#include <sstream>

int day2::password(std::vector<std::string> in) {
  int result{0};
  for (const auto& str : in) {
    int from, to;
    char c, policy;
    std::string pw;
    std::stringstream stream{str};
    stream >> from;
    stream.get(c);
    stream >> to;
    stream.get(c);
    stream.get(policy);
    stream.get(c);
    stream >> pw;

    int countPolicy{0};
    for (char c : pw) {
      if (c == policy) {
        countPolicy++;
      }
    }
    if (countPolicy >= from && countPolicy <= to) {
      result++;
    }
  }
  return result;
}

int day2::password2(std::vector<std::string> in) {
  int result{0};
  for (const auto& str : in) {
    int from, to;
    char c, policy;
    std::string pw;
    std::stringstream stream{str};
    stream >> from;
    stream.get(c);
    stream >> to;
    stream.get(c);
    stream.get(policy);
    stream.get(c);
    stream >> pw;

    if ((pw[from - 1] == policy || pw[to - 1] == policy) &&
        pw[from - 1] != pw[to - 1]) {
      result++;
    }
  }
  return result;
}
