#include "day11.h"

#include "helpers/helpers.h"

day11::Plan day11::parsePlan(const std::vector<std::string>& in) {
  Plan result;
  for (const auto& line : in) {
    auto& current = result.emplace_back();
    for (char ch : line) {
      switch (ch) {
        case '.':
          current.push_back(Seat::Floor);
          break;
        case '#':
          current.push_back(Seat::Occupied);
          break;
        case 'L':
          current.push_back(Seat::Empty);
          break;
      }
    }
  }
  return result;
}

std::pair<day11::Plan, bool> day11::doStep(
    const Plan& plan, std::function<Seat(const Plan&, size_t i, size_t j)> f) {
  auto result = std::make_pair<Plan, bool>({}, false);
  for (size_t i = 0; i < plan.size(); ++i) {
    auto& current = result.first.emplace_back();
    for (size_t j = 0; j < plan[i].size(); ++j) {
      Seat s = f(plan, i, j);
      if (s != plan[i][j]) {
        result.second = true;
      }
      current.push_back(s);
    }
  }
  return result;
}

day11::Seat day11::newSeat(const Plan& plan, size_t i, size_t j) {
  auto check = [&plan](size_t i, int di, size_t j, int dj) {
    if (i + di < plan.size() && j + dj < plan[i].size()) {
      return plan[i + di][j + dj] == Seat::Occupied;
    }
    return false;
  };
  int count{0};
  if (check(i, -1, j, -1)) count++;
  if (check(i, -1, j, 0)) count++;
  if (check(i, -1, j, 1)) count++;
  if (check(i, 0, j, -1)) count++;
  if (check(i, 0, j, 1)) count++;
  if (check(i, 1, j, -1)) count++;
  if (check(i, 1, j, 0)) count++;
  if (check(i, 1, j, 1)) count++;
  if (plan[i][j] == Seat::Occupied) {
    if (count >= 4) {
      return Seat::Empty;
    }
  } else if (plan[i][j] == Seat::Empty) {
    if (count == 0) {
      return Seat::Occupied;
    }
  }
  return plan[i][j];
}

int64_t day11::puzzle1(const std::vector<std::string>& in) {
  auto plan = parsePlan(in);
  bool changed = true;
  while (changed) {
    auto res = doStep(plan, newSeat);
    changed = res.second;
    plan = std::move(res.first);
  }

  int count{0};
  for (const auto& line : plan) {
    for (const auto seat : line) {
      if (seat == Seat::Occupied) {
        ++count;
      }
    }
  }
  return count;
}

void day11::print(const Plan& plan) {
  std::cout << "\n";
  for (const auto& line : plan) {
    for (const auto seat : line) {
      switch (seat) {
        case Seat::Occupied:
          std::cout << "#";
          break;
        case Seat::Empty:
          std::cout << "L";
          break;
        case Seat::Floor:
          std::cout << ".";
          break;
      }
    }
    std::cout << "\n";
  }
}

day11::Seat day11::newSeat2(const Plan& plan, size_t i, size_t j) {
  auto check = [&plan](size_t i, int di, size_t j, int dj) {
    while (i + di < plan.size() && j + dj < plan[i].size()) {
      i = i + di;
      j = j + dj;
      if (plan[i][j] == Seat::Floor) continue;
      return plan[i][j] == Seat::Occupied;
    }
    return false;
  };
  int count{0};
  if (check(i, -1, j, -1)) count++;
  if (check(i, -1, j, 0)) count++;
  if (check(i, -1, j, 1)) count++;
  if (check(i, 0, j, -1)) count++;
  if (check(i, 0, j, 1)) count++;
  if (check(i, 1, j, -1)) count++;
  if (check(i, 1, j, 0)) count++;
  if (check(i, 1, j, 1)) count++;
  if (plan[i][j] == Seat::Occupied) {
    if (count >= 5) {
      return Seat::Empty;
    }
  } else if (plan[i][j] == Seat::Empty) {
    if (count == 0) {
      return Seat::Occupied;
    }
  }
  return plan[i][j];
}

int64_t day11::puzzle2(const std::vector<std::string>& in) {
  auto plan = parsePlan(in);
  bool changed = true;
  while (changed) {
    auto res = doStep(plan, newSeat2);
    changed = res.second;
    plan = std::move(res.first);
  }

  int count{0};
  for (const auto& line : plan) {
    for (const auto seat : line) {
      if (seat == Seat::Occupied) {
        ++count;
      }
    }
  }
  return count;
}