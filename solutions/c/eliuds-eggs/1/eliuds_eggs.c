#include "eliuds_eggs.h"

int egg_count(int count) {
    int eggs = 0;
    int c = count;
    while (c > 0) {
        if (c % 2 == 1) {
            c /= 2;
            eggs += 1;
        } else {
            c /= 2;
        }
    }
    return eggs;
} 