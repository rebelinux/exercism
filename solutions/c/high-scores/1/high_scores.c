#include "high_scores.h"

int32_t latest(const int32_t *scores, size_t scores_len) {
    return scores[scores_len - 1];
}

int32_t personal_best(const int32_t *scores, size_t scores_len) {
    int32_t highest = 0;
    for (uint32_t i = 0; i < scores_len; i++) {
        if (scores[i] > highest) {
            highest = scores[i];
        }
    }
    return highest;
}

size_t personal_top_three(const int32_t *scores, size_t scores_len, int32_t *output) {
    int32_t highest1 = INT32_MIN;
    int32_t highest2 = INT32_MIN;
    int32_t highest3 = INT32_MIN;
    size_t arrlen = 0;

    for (uint32_t i = 0; i < scores_len; i++) {
        if (scores[i] > highest1) {
            highest3 = highest2;
            highest2 = highest1;
            highest1 = scores[i];
        } else if (scores[i] > highest2) {
            highest3 = highest2;
            highest2 = scores[i];
        } else if (scores[i] > highest3) { 
            highest3 = scores[i];
        } 
    }
    if (highest1 != INT32_MIN) {
        output[arrlen++] = highest1;
    }
    if (highest2 != INT32_MIN) {
        output[arrlen++] = highest2;
    }
    if (highest3 != INT32_MIN) {
        output[arrlen++] = highest3;
    }

    return arrlen;
}