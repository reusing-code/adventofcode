#pragma once

#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

namespace helpers {
using StringGroups = std::vector<std::vector<std::string>>;

StringGroups splitByEmptyLine(const std::vector<std::string>& in);
std::vector<std::string> split(const std::string& in, char separator);
}  // namespace helpers