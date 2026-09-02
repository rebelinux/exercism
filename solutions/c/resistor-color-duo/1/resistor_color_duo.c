#include "resistor_color_duo.h"
#include <stdio.h>

int color_code(resistor_band_t *a) {
    int value[2];
    for (int i = 0; i < 2; i++) {
        value[i] = a[i];
    }
    return value[0] * 10 + value[1];
}

// int main(void) {
//     // resistor_band_t a = {BLACK, BROWN};
//     int b = color_code((resistor_band_t[]){ BROWN, BLACK });
//     printf("%d\n", b);
//     return 0;
// }