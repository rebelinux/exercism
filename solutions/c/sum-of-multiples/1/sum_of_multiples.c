#include "sum_of_multiples.h"
#include "stdbool.h"

bool validate(int *array, size_t size, int number) {
    bool result = false;
    for (size_t i = 0; i < size; i++)
    {
        if (array[i] == number) {
            result = true;
            break; 
        }
    }
    return result;
}

unsigned int sum(const unsigned int *factors, const size_t number_of_factors, const unsigned int limit) {
    unsigned int sum = 0;
    int array[100000] = {0};
    int index = 0;
    for (size_t i = 0; i < number_of_factors; i++) {
        if (factors[i] == 0) {break;} 
        unsigned int s = 0;
        while (s < limit)
        {
            if (!validate(array, sizeof(array) / sizeof(array[0]), s)) {
                array[index] = s;
                sum += s;
            }
            s += factors[i];    
            index++;
        }
    }
    return sum;
}