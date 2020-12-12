#include "day8.h"

#include <sstream>

#include "helpers/helpers.h"

int day8::puzzle1(const std::vector<std::string>& in) {
  Handheld hh{in};
  hh.run();
  return hh.accumulator();
}

int day8::puzzle2(const std::vector<std::string>& in) {
  Handheld original{in};
  for (int i = 0; i < original.program().size(); ++i) {
    Handheld copy = original;
    auto& instr = copy.program()[i];
    switch (instr.op) {
      case Instruction::Operation::acc:
        continue;
      case Instruction::Operation::jmp:
        instr.op = Instruction::Operation::nop;
        break;
      case Instruction::Operation::nop:
        instr.op = Instruction::Operation::jmp;
        break;
    }
    if (copy.run() == Handheld::Status::ok) {
      return copy.accumulator();
    }
  }

  return 0;
}

void Handheld::parseProgram(const std::vector<std::string>& in) {
  m_program.clear();
  for (const auto& instr : in) {
    m_program.push_back(parseInstruction(instr));
  }
}

Handheld::Handheld(const std::vector<std::string>& in) { parseProgram(in); }

Handheld::Status Handheld::run(int pc) {
  m_accumulator = 0;
  while (m_visited.find(pc) == m_visited.end()) {
    if (pc == m_program.size()) {
      return Status::ok;
    } else if (pc > m_program.size()) {
      return Status::corrupted;
    }
    m_visited.insert(pc);
    const auto& instr = m_program[pc];
    switch (instr.op) {
      case Instruction::Operation::acc:
        m_accumulator += instr.param;
        ++pc;
        break;
      case Instruction::Operation::jmp:
        pc += instr.param;
        break;
      case Instruction::Operation::nop:
        ++pc;
        break;
    }
  }

  return Status::corrupted;
}

int Handheld::accumulator() const { return m_accumulator; }

Handheld::Program& Handheld::program() { return m_program; }

Instruction Handheld::parseInstruction(const std::string& str) {
  std::stringstream stream{str};
  std::string op;
  int param{0};
  stream >> op >> param;
  return {Instruction::parseOperation(op), param};
}

Instruction::Operation Instruction::parseOperation(const std::string& str) {
  if (str == "acc") {
    return Operation::acc;
  }
  if (str == "jmp") {
    return Operation::jmp;
  }
  return Operation::nop;
}