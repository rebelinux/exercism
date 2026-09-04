#include "hamming.h"
#include "stdio.h"
#include "string.h"

int compute(const char *lhs, const char *rhs) {
    int r = 0;
    if (strlen(rhs) == 0 && strlen(lhs) == 0) {
        return 0;
    } else if (strlen(lhs) != strlen(rhs) || strlen(rhs) == 0 || strlen(lhs) == 0) {
        return -1;
    } else if (strlen(lhs) == 1)
    {
        if (lhs != rhs) {
            return 1;
        } else {
            return 0;
        } 
    }
    
    for (size_t i = 0; i < strlen(lhs); i++)
    {
        if (lhs[i] != rhs[i])
        {
            r++;
        }
    } 
    return r;  
}