#pragma once

#include <iostream>
#include <string>
#include <unordered_set>
#include <vector>

struct Instruction {
  enum class Operation { acc, jmp, nop };
  Operation op{Operation::nop};
  int param{0};

  static Operation parseOperation(const std::string& str);
};

class Handheld {
 public:
  using Program = std::vector<Instruction>;
  enum class Status { ok, corrupted };
  Handheld(const std::vector<std::string>& in);

  Status run(int startPC = 0);
  int accumulator() const;
  Program& program();

 private:
  void parseProgram(const std::vector<std::string>& in);
  static Instruction parseInstruction(const std::string& str);

  Program m_program;
  int m_accumulator{0};
  std::unordered_set<int> m_visited;
};

class day8 {
 public:
  static int puzzle1(const std::vector<std::string>& in);
  static int puzzle2(const std::vector<std::string>& in);

 private:
};