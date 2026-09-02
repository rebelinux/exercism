#include "resistor_color_duo.h"

int color_code(resistor_band_t *a) {
    return a[0] * 10 + a[1];
}