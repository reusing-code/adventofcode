#include "day13.h"

#include <algorithm>
#include <execution>

#include "helpers/helpers.h"

int64_t day13::puzzle1(const std::vector<std::string>& in) {
  int64_t startTime = std::stoll(in[0]);
  std::vector<int64_t> busses;
  auto splitted = helpers::split(in[1], ',');
  for (const auto& s : splitted) {
    if (s != "x") {
      busses.push_back(std::stoll(s));
    }
  }

  int64_t minTime = std::numeric_limits<int64_t>::max();
  int64_t minBus = 0;

  for (int64_t bus : busses) {
    int64_t factor = startTime / bus;
    if (startTime % bus != 0) {
      ++factor;
    }
    int64_t time = factor * bus;
    if (time < minTime) {
      minTime = time;
      minBus = bus;
    }
  }

  return (minTime - startTime) * minBus;
}

int64_t day13::puzzle2(const std::vector<std::string>& in) {
  std::vector<std::pair<int64_t, int64_t>> busses;
  auto splitted = helpers::split(in[1], ',');
  for (size_t i = 0; i < splitted.size(); ++i) {
    if (splitted[i] != "x") {
      busses.emplace_back(std::stoll(splitted[i]), i);
    }
  }

  std::sort(busses.begin(), busses.end(),
            [](const auto& a, const auto& b) { return a.first > b.first; });

  for (int64_t i = busses[0].first;; i += busses[0].first) {
    if (std::all_of(std::execution::par, busses.begin() + 1, busses.end(),
                    [ii = i - busses[0].second](const auto& bus) {
                      return (ii + bus.second) % bus.first == 0;
                    })) {
      return i - busses[0].second;
    }
  }

  return 0;
}