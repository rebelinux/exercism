#include "armstrong_numbers.h"
#include "math.h"


bool is_armstrong_number(int candidate) {
    if (candidate < 0) {
        return false;
    }
    
    int value = candidate;
    int num[20];
    int index = 0;
    int sum = 0;
    
    do
    {
        num[index] = value % 10;
        value = value / 10;
        index++;
    } while (value != 0);

    for (int i = 0; i < index; i++) {
        sum = sum + (int)pow(num[i], index);
    }

    return sum == candidate;
}