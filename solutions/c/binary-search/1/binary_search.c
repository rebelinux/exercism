#include "binary_search.h"
#include "stdbool.h"


const int *binary_search(int value, const int *arr, size_t length) {
    const int middle = arr[length / 2];
    bool match = false;
    while (!match)
    {
        if (value >= middle) {
            for (int i = length / 2; i < (int)length; i++)
            {
                if (arr[i] == value)
                {
                    match = true;
                    return &arr[i];
                }   
            } 
            return NULL;
        } else if (value <= middle) {
            for (int i = length / 2; i >= 0; i--)
            {
                if (arr[i] == value)
                {
                    match = true;
                    return &arr[i];
                }
            }
            return NULL;
        }
    }
    return NULL;
}