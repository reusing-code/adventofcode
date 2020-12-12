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

void splitRecursive(const std::string& in, char separator,
                    std::vector<std::string>& out) {
  auto find = in.find(separator);
  out.push_back(in.substr(0, find));
  if (find == std::string::npos) {
    return;
  }
  splitRecursive(in.substr(find + 1), separator, out);
}

std::vector<std::string> split(const std::string& in, char separator) {
  std::vector<std::string> result;
  splitRecursive(in, separator, result);
  return result;
}

std::vector<int64_t> parseIntVec(const std::vector<std::string>& in) {
  std::vector<int64_t> result;
  result.reserve(in.size());
  std::transform(in.begin(), in.end(), std::back_inserter(result),
                 [](const std::string& str) { return std::stoll(str); });
  return result;
}
}  // namespace helpers