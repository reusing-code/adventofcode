#include <string>
#include <vector>

class FFT {
 public:
  explicit FFT(const std::vector<char>& basePattern);

  std::string run(const std::string& input, uint32_t iterations);
  std::string runPartTwo(const std::string& input, uint32_t iterations);

 private:
  void iteration();
  void iteration2();
  void createPatterns(uint32_t length);
  int getSummand(int row, int column);
  std::vector<std::vector<char>> patterns_;
  std::vector<char> basePattern_;
  std::vector<char> currentValue_;
  size_t offset_{0};
};