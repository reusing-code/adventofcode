#include "day16.h"

#include <algorithm>
#include <set>
#include <unordered_map>

#include "helpers/helpers.h"

int64_t day16::puzzle1(const std::vector<std::string>& in) {
  auto ts = TicketSystem::parseTicketSystem(in);
  return ts.errorRate();
}

int64_t day16::puzzle2(const std::vector<std::string>& in) {
  auto ts = TicketSystem::parseTicketSystem(in);
  return ts.departure();
}

day16::TicketSystem day16::TicketSystem::parseTicketSystem(
    const std::vector<std::string>& in) {
  TicketSystem result;
  auto groups = helpers::splitByEmptyLine(in);
  for (const auto& str : groups[0]) {
    TicketSystem::Rule rule;
    auto split = helpers::split(str, ':');
    rule.name = split[0];
    split = helpers::split(split[1], 'o');
    auto parseRange = [](const std::string& in, Range& range) {
      auto split = helpers::split(in, '-');
      range.from = std::stoll(split[0]);
      range.to = std::stoll(split[1]);
    };
    parseRange(split[0], rule.range1);
    parseRange(split[1].substr(1), rule.range2);
    result.rules.emplace_back(std::move(rule));
  }
  result.ticket = helpers::splitInt(groups[1][1], ',');
  for (size_t i = 1; i < groups[2].size(); ++i) {
    result.nearbyTickets.emplace_back(helpers::splitInt(groups[2][i], ','));
  }
  return result;
}

int64_t day16::TicketSystem::errorRate() {
  int64_t result{0};
  std::vector<Range> ranges;
  for (const auto& rule : rules) {
    ranges.push_back(rule.range1);
    ranges.push_back(rule.range2);
  }
  auto checkRanges = [&ranges](int64_t val) {
    for (const auto& range : ranges) {
      if (val >= range.from && val <= range.to) {
        return true;
      }
    }
    return false;
  };

  for (const auto& tick : nearbyTickets) {
    for (int64_t val : tick) {
      if (!checkRanges(val)) {
        result += val;
      }
    }
  }

  return result;
}

int64_t day16::TicketSystem::departure() {
  int64_t result{1};
  std::vector<Range> ranges;
  for (const auto& rule : rules) {
    ranges.push_back(rule.range1);
    ranges.push_back(rule.range2);
  }
  auto checkRanges = [&ranges](int64_t val) {
    for (const auto& range : ranges) {
      if (val >= range.from && val <= range.to) {
        return true;
      }
    }
    return false;
  };

  auto checkTicket = [checkRanges](const auto& tick) {
    for (int64_t val : tick) {
      if (!checkRanges(val)) {
        return false;
      }
    }
    return true;
  };
  std::vector<Ticket> validTickets;
  for (const auto& tick : nearbyTickets) {
    if (checkTicket(tick)) {
      validTickets.push_back(tick);
    }
  }
  std::vector<std::set<int64_t>> categoryValues(ticket.size());
  for (const auto& tick : validTickets) {
    for (size_t i = 0; i < tick.size(); ++i) {
      categoryValues[i].insert(tick[i]);
    }
  }

  std::unordered_map<int64_t, std::set<int64_t>> possibleCategories;
  for (size_t i = 0; i < ticket.size(); ++i) {
    const auto& values = categoryValues[i];
    for (size_t j = 0; j < rules.size(); ++j) {
      const auto& cat = rules[j];
      auto minmax = std::minmax_element(values.begin(), values.end());
      if (*minmax.first < cat.range1.from || *minmax.second > cat.range2.to) {
        continue;
      }
      if (std::any_of(values.begin(), values.end(), [&cat](int64_t a) {
            return a > cat.range1.to && a < cat.range2.from;
          })) {
        continue;
      }
      possibleCategories[i].insert(j);
    }
  }
  std::vector<int64_t> correctCategories(ticket.size());
  while (!possibleCategories.empty()) {
    auto elem =
        std::find_if(possibleCategories.begin(), possibleCategories.end(),
                     [](const auto& a) { return a.second.size() == 1; });
    int64_t cat = *(elem->second.begin());
    correctCategories[elem->first] = cat;
    for (auto& it : possibleCategories) {
      it.second.erase(cat);
    }
    possibleCategories.erase(elem->first);
  }

  for (size_t i = 0; i < correctCategories.size(); ++i) {
    int64_t rule = correctCategories[i];
    if (rules[rule].name.find("departure") == 0) {
      result *= ticket[i];
    }
  }

  return result;
}