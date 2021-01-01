#include "day18.h"

#include <algorithm>
#include <numeric>
#include <stack>

#include "helpers/helpers.h"

int64_t day18::puzzle1(const std::vector<std::string>& in) {
  return std::accumulate(
      in.begin(), in.end(), 0ll,
      [](int64_t a, const std::string& b) { return a + solveEquation(b); });
}

int64_t day18::puzzle2(const std::vector<std::string>& in) {
  return std::accumulate(
      in.begin(), in.end(), 0ll,
      [](int64_t a, const std::string& b) { return a + solveEquation2(b); });
}

int64_t day18::solveEquation(const std::string& in) {
  std::stack<std::pair<int64_t, Op>> stack;
  stack.emplace(0, Op::NONE);

  auto consumeValue = [&stack](int64_t val) {
    auto& top = stack.top();
    switch (top.second) {
      case Op::NONE:
        top.first = val;
        break;
      case Op::ADD:
        top.first = top.first + val;
        top.second = Op::NONE;
        break;
      case Op::MUL:
        top.first = top.first * val;
        top.second = Op::NONE;
        break;
    }
  };

  for (char c : in) {
    if (c >= '0' && c <= '9') {
      int64_t val = c - '0';
      consumeValue(val);
    } else if (c == '+') {
      stack.top().second = Op::ADD;
    } else if (c == '*') {
      stack.top().second = Op::MUL;
    } else if (c == '(') {
      stack.emplace(0, Op::NONE);
    } else if (c == ')') {
      int64_t val = stack.top().first;
      stack.pop();
      consumeValue(val);
    }
  }
  return stack.top().first;
}

int64_t day18::solveEquation2(const std::string& in) {
  std::stack<std::vector<std::pair<int64_t, Op>>> stack;
  stack.emplace();
  stack.top().emplace_back(0, Op::NONE);

  auto evalTop = [&stack]() -> int64_t {
    auto& top = stack.top();
    std::vector<std::pair<int64_t, Op>> mult;
    for (const auto& elem : top) {
      if (elem.second == Op::ADD) {
        mult.back().first += elem.first;
      } else {
        mult.push_back(elem);
      }
    }
    int64_t result{0};
    for (const auto& elem : mult) {
      if (elem.second == Op::NONE) {
        result = elem.first;
      } else if (elem.second == Op::MUL) {
        result *= elem.first;
      }
    }
    return result;
  };

  for (char c : in) {
    if (c >= '0' && c <= '9') {
      int64_t val = c - '0';
      stack.top().back().first = val;
    } else if (c == '+') {
      stack.top().emplace_back(0, Op::ADD);
    } else if (c == '*') {
      stack.top().emplace_back(0, Op::MUL);
    } else if (c == '(') {
      stack.emplace();
      stack.top().emplace_back(0, Op::NONE);
    } else if (c == ')') {
      int64_t val = evalTop();
      stack.pop();
      stack.top().back().first = val;
    }
  }
  return evalTop();
}