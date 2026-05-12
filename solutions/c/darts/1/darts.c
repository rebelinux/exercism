#include "darts.h"
#include <math.h>

int score(coordinate_t landing_position) {
    float distance = sqrt(landing_position.x * landing_position.x + landing_position.y * landing_position.y);
    
    if (distance <= 1.0f) {
        return 10;
    } else if (distance <= 5.0f) {
        return 5;
    } else if (distance <= 10.0f) {
        return 1;
    }
    return 0;
}