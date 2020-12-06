#pragma once

#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

namespace helpers {
using StringGroups = std::vector<std::vector<std::string>>;

StringGroups splitByEmptyLine(const std::vector<std::string>& in);
}  // namespace helpers