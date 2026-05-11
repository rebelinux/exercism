#include "collatz_conjecture.h"

int steps(int start) {
    if (start <= 0) {return ERROR_VALUE;}
    int end = 1;
    int cur = start;
    int index = 0;
    while (cur != end)
    {
        if (cur % 2 == 0) {
            cur /= 2;
            index++;
        } else {
            cur = (cur * 3) + 1;
            index++;
        }

    }
    return index;
}