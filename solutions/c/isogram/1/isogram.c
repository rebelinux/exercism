#include "isogram.h"
#include "string.h" 
#include <ctype.h>

bool is_isogram(const char phrase[]) {
    if (phrase == NULL) {return false;}
    char currentchar = tolower(phrase[0]);
    bool value = true;
    for (size_t i = 1; i < strlen(phrase); i++) {
        for (size_t ii = i; ii < strlen(phrase); ii++) {
            if (tolower(phrase[ii]) == '-') {
                continue;
            } if (tolower(phrase[ii]) == ' ') {
                continue;
            } else if (tolower(phrase[ii]) == currentchar) {
                value = false;
                break;
            }
        }
        currentchar = tolower(phrase[i]);
    }
    return value;
}