#include "helpers.h"

namespace helpers {
StringGroups splitByEmptyLine(const std::vector<std::string>& in) {
  StringGroups result;
  auto it = in.begin();
  while (it != in.end()) {
    auto it2 = std::find(it, in.end(), "");
    result.emplace_back(it, it2);
    it = it2 == in.end() ? in.end() : it2 + 1;
  }
  return result;
}
}  // namespace helpers