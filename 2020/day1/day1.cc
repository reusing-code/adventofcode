#include "day1.h"

#include <algorithm>
#include <iostream>

uint64_t day1::expense(std::vector<uint64_t> values) {
    std::sort(values.begin(), values.end());
    auto head = values.begin();
    auto tail = values.end() - 1;
    while (head < tail) {
        auto res = *head + *tail;
        if (res ==2020) {
            return *head * *tail;
        }
        if ( res > 2020) {
            tail--;
        } else {
            head++;
        }
    }
    return 1;
}

uint64_t day1::expense3(std::vector<uint64_t> values) {
    std::sort(values.begin(), values.end());
    for (auto it: values) {
    auto head = values.begin();
    auto tail = values.end() - 1;
    while (head < tail) {
        if (it == *head) {
            head++;
            continue;
        }
        if (it == *tail) {
            tail--;
            continue;
        }
        auto res = *head + *tail + it;
        if (res ==2020) {
            return *head * *tail * it;
        }
        if ( res > 2020) {
            tail--;
        } else {
            head++;
        }
    }
    }
    return 1;
}

