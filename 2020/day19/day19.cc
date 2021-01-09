#include "day19.h"

#include <bits/stdint-uintn.h>

#include <algorithm>
#include <memory>
#include <queue>
#include <sstream>
#include <string>
#include <unordered_map>
#include <unordered_set>

#include "helpers/helpers.h"

const std::unordered_set<std::string>& getPossibleStrings(
    std::shared_ptr<std::unordered_map<uint64_t, std::string>> rulesRaw,
    std::shared_ptr<
        std::unordered_map<uint64_t, std::unordered_set<std::string>>>
        cache,
    uint64_t start) {
  std::queue<uint64_t> queue;
  queue.push(start);
  while (!queue.empty()) {
    auto id = queue.front();
    queue.pop();

    if (cache->find(id) != cache->end()) {
      continue;
    }
    std::unordered_set<std::string> result;
    auto& rule = rulesRaw->at(id);
    auto findQuote = rule.find("\"");
    if (findQuote != std::string::npos) {
      result.insert(rule.substr(findQuote + 1, 1));
      (*cache)[id] = result;
      continue;
    }
    auto parse = [&cache, &result, &queue](const std::string& str) -> bool {
      bool found = true;
      std::stringstream stream{str};
      uint64_t first, second;
      stream >> first >> second;
      if (stream.fail()) {
        if (cache->contains(first)) {
          auto& c = (*cache)[first];
          for (const auto& a : c) {
            result.insert(a);
          }
        } else {
          found = false;
          queue.push(first);
        }
      } else {
        if (cache->contains(first) && cache->contains(second)) {
          auto& c1 = (*cache)[first];
          auto& c2 = (*cache)[second];
          for (const auto& a : c1) {
            for (const auto& b : c2) {
              result.insert(a + b);
            }
          }
        } else {
          found = false;
          queue.push(first);
          queue.push(second);
        }
      }
      return found;
    };
    bool foundAll = true;
    foundAll &= parse(rule);

    auto findAlt = rule.find("|");
    if (findAlt != std::string::npos) {
      foundAll &= parse(rule.substr(findAlt + 1));
    }
    if (foundAll) {
      (*cache)[id] = result;
    } else {
      queue.push(id);
    }
  }
  return cache->find(start)->second;
}

int64_t day19::puzzle1(const std::vector<std::string>& in) {
  auto inGroups = helpers::splitByEmptyLine(in);
  const auto& rulesIn = inGroups.at(0);
  const auto& messages = inGroups.at(1);
  auto rulesRaw = std::make_shared<std::unordered_map<uint64_t, std::string>>();
  for (const auto& rule : rulesIn) {
    auto colon = rule.find(":");
    uint64_t id = std::stoull(rule.substr(0, colon));
    (*rulesRaw)[id] = rule.substr(colon + 2);
  }
  auto cache = std::make_shared<
      std::unordered_map<uint64_t, std::unordered_set<std::string>>>();
  std::queue<uint64_t> queue;
  auto possibleStrings = getPossibleStrings(rulesRaw, cache, 0);

  int64_t result{0};
  for (const auto& message : messages) {
    if (possibleStrings.contains(message)) {
      result++;
    }
  }
  return result;
}

int64_t day19::puzzle2(const std::vector<std::string>& in) {
  (void)in;
  return 0;
}