#pragma once

#include <iostream>
#include <string>
#include <vector>

class day16 {
 public:
  static int64_t puzzle1(const std::vector<std::string>& in);
  static int64_t puzzle2(const std::vector<std::string>& in);

  struct TicketSystem {
    using Ticket = std::vector<int64_t>;
    Ticket ticket;
    std::vector<Ticket> nearbyTickets;
    struct Range {
      int64_t from{0};
      int64_t to{0};
    };
    struct Rule {
      std::string name;

      Range range1;
      Range range2;
    };
    std::vector<Rule> rules;

    static TicketSystem parseTicketSystem(const std::vector<std::string>& in);
    int64_t errorRate();
    int64_t departure();
  };

 private:
};