#include "day7.h"

#include <queue>
#include <sstream>
#include <unordered_map>
#include <unordered_set>

#include "2020/helpers/helpers.h"

const std::string CONTAINS_STR{" bags contain "};
const std::string NO_BAGS{"no other bags."};
const std::string BAG_STR{" bag"};

std::pair<std::string, std::vector<std::pair<std::string, int>>> parseRule(
    const std::string& in) {
  std::pair<std::string, std::vector<std::pair<std::string, int>>> result;
  auto contains = in.find(CONTAINS_STR);
  if (contains != std::string::npos) {
    auto outerBag = in.substr(0, contains);
    result.first = outerBag;
    auto innerBags = in.substr(contains + CONTAINS_STR.size());
    if (innerBags != "no other bags.") {
      auto splitBags = helpers::split(innerBags, ',');
      for (const auto& bag : splitBags) {
        std::stringstream stream{bag};
        int count{0};
        stream >> count;
        std::string rest;
        std::getline(stream, rest);
        auto findBags = rest.find(BAG_STR);
        result.second.emplace_back(rest.substr(1, findBags - 1), count);
      }
    }
  }
  return result;
}

int day7::puzzle1(const std::vector<std::string>& in) {
  std::unordered_map<std::string, std::vector<std::string>> tree;
  for (const auto& str : in) {
    auto res = parseRule(str);
    for (const auto& it : res.second) {
      tree[it.first].push_back(res.first);
    }
  }
  std::unordered_set<std::string> bags;
  std::queue<std::string> queue;
  queue.push("shiny gold");
  while (!queue.empty()) {
    for (const auto& it : tree[queue.front()]) {
      if (bags.find(it) == bags.end()) {
        bags.insert(it);
        queue.push(it);
      }
    }
    queue.pop();
  }
  return bags.size();
}

int countBagsRec(
    const std::unordered_map<std::string,
                             std::vector<std::pair<std::string, int>>>& tree,
    const std::string& in) {
  int result{0};
  for (const auto& bag : tree.at(in)) {
    result += bag.second + bag.second * countBagsRec(tree, bag.first);
  }
  return result;
}

int day7::puzzle2(const std::vector<std::string>& in) {
  std::unordered_map<std::string, std::vector<std::pair<std::string, int>>>
      tree;
  for (const auto& str : in) {
    auto res = parseRule(str);
    auto& treeElem = tree[res.first];
    for (const auto& it : res.second) {
      treeElem.push_back(it);
    }
  }
  return countBagsRec(tree, "shiny gold");
}