#pragma once

#include <functional>
#include <iostream>
#include <string>
#include <vector>

class day11 {
 public:
  enum class Seat { Floor, Empty, Occupied };

  using Plan = std::vector<std::vector<Seat>>;

  static Plan parsePlan(const std::vector<std::string>& in);
  static std::pair<Plan, bool> doStep(
      const Plan& plan, std::function<Seat(const Plan&, size_t i, size_t j)> f);
  static Seat newSeat(const Plan& plan, size_t i, size_t j);
  static Seat newSeat2(const Plan& plan, size_t i, size_t j);
  static void print(const Plan& plan);
  static int64_t puzzle1(const std::vector<std::string>& in);
  static int64_t puzzle2(const std::vector<std::string>& in);

 private:
};