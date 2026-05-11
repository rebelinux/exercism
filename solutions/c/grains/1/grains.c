#include "grains.h"

uint64_t square(uint8_t index) {
    uint64_t number = 1;
    if (index == 0) {
        return 0;
    }
    for (int i = 1; i < index; i++) {
        number *= 2;
    }
    return number;
}

uint64_t total(void) {
    uint64_t sum = 0;
    for (int i = 1; i <= 64; i++) {
        sum += square(i);
    }
    return sum;
}