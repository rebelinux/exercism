#include "luhn.h"
#include "string.h"
#include "stdlib.h"


bool luhn(const char *num) {
    bool value;
    size_t len = strlen(num);
    int ch;
    if (len <= 1) {
        value =  false;
    }
    for (int i = len - 2; i >= 0; i--) {
        int temp = num[i] - '0';
        if (temp >= 0 && temp <= 9) {
            if ((temp + temp) > 9) {
                ch += (temp + temp) - 9;
            } else {
                ch += temp;
            }
        } else {
            continue;
        }
    }
    ch += (num[len - 1] - '0');
    
    if (ch % 10 == '0') {
        value = true;
    }

    return value;
}