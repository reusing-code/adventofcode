#include "day12.h"

#include "helpers/helpers.h"

int64_t day12::puzzle1(const std::vector<std::string>& in) {
  Nav nav;
  for (const auto& str : in) {
    nav.doStep(str);
  }
  return std::abs(nav.x) + std::abs(nav.y);
}

int64_t day12::puzzle2(const std::vector<std::string>& in) {
  Nav2 nav;
  for (const auto& str : in) {
    nav.doStep(str);
  }
  return std::abs(nav.ship.x) + std::abs(nav.ship.y);
}

void day12::Nav::doStep(const std::string& instr) {
  char c = instr[0];
  int64_t param = std::stoll(instr.substr(1));

  if (c == 'F') {
    if (dir == 0) c = 'N';
    if (dir == 1) c = 'E';
    if (dir == 2) c = 'S';
    if (dir == 3) c = 'W';
  }

  switch (c) {
    case 'N':
      y += param;
      break;
    case 'S':
      y -= param;
      break;
    case 'E':
      x += param;
      break;
    case 'W':
      x -= param;
      break;
    case 'L':
      dir = (dir - (param / 90) + 4) % 4;
      break;
    case 'R':
      dir = (dir + (param / 90)) % 4;
      break;
  }
}

void day12::Nav2::doStep(const std::string& instr) {
  char c = instr[0];
  int64_t param = std::stoll(instr.substr(1));

  switch (c) {
    case 'N':
      waypoint.y += param;
      break;
    case 'S':
      waypoint.y -= param;
      break;
    case 'E':
      waypoint.x += param;
      break;
    case 'W':
      waypoint.x -= param;
      break;
    case 'L': {
      int dir = (-(param / 90) + 4) % 4;
      rotate(dir);
      break;
    }
    case 'R': {
      int dir = ((param / 90)) % 4;
      rotate(dir);
      break;
    }
    case 'F':
      ship.x += param * waypoint.x;
      ship.y += param * waypoint.y;
      break;
  }
}

void day12::Nav2::rotate(int dir) {
  switch (dir) {
    case 0:
      break;
    case 1:
      std::swap(waypoint.x, waypoint.y);
      waypoint.y *= -1;
      break;
    case 2:
      waypoint.x *= -1;
      waypoint.y *= -1;
      break;
    case 3:
      std::swap(waypoint.x, waypoint.y);
      waypoint.x *= -1;
      break;
  }
}
