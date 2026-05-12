#include "binary.h"
#include "string.h"
#include "math.h"

int convert(const char *input) {
    int index = 0;
    int result = 0;
    size_t len = strlen(input);
    for (int i = len -1; i >= 0; i--) {
        int value = 0;
        if (input[i] == '1') {
            value = pow(2, index);
            index++;
        } else if (input[i] == '0') {
            index++;
        } else {
            return INVALID;
        }
        result += value;
    }
    return result;
}