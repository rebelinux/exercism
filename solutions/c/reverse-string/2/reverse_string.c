#include "reverse_string.h"
#include "string.h"
#include "stdlib.h"

char *reverse(const char *value) {
    int len = strlen(value);
    char *str = malloc(len + 1);
    for (int i = len; i >= 0; i--)
    {
        str[len -1 - i] = value[i];
    }
    str[len] = '\0';
    return str;
}
