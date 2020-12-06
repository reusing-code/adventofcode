#include "day6.h"

#include <iterator>
#include <set>
#include <unordered_set>

#include "2020/helpers/helpers.h"

int day6::puzzle1(const std::vector<std::string>& in) {
  int result{0};
  auto groups = helpers::splitByEmptyLine(in);
  for (const auto& group : groups) {
    std::unordered_set<char> questions;
    for (const auto& person : group) {
      for (char c : person) {
        questions.insert(c);
      }
    }
    result += questions.size();
  }
  return result;
}

int day6::puzzle2(const std::vector<std::string>& in) {
  int result{0};
  auto groups = helpers::splitByEmptyLine(in);
  for (const auto& group : groups) {
    bool first{true};
    std::vector<char> questions;
    for (const auto& person : group) {
      std::set<char> personQuestions;
      for (char c : person) {
        personQuestions.insert(c);
      }
      if (first) {
        std::copy(personQuestions.begin(), personQuestions.end(),
                  std::back_inserter(questions));
        first = false;
      } else {
        std::vector<char> newQuestions;
        std::set_intersection(questions.begin(), questions.end(),
                              personQuestions.begin(), personQuestions.end(),
                              std::back_inserter(newQuestions));
        questions = std::move(newQuestions);
      }
    }
    result += questions.size();
  }
  return result;
}