#include "space_age.h"
#include "stdio.h"
#include "math.h"

double sec_to_years(double planetinyears, int64_t secs) {
    return secs / ((double)round(365.25 * planetinyears) * 86400);
}

float age(planet_t planet, int64_t seconds) {
    double planet_years[] = {0.2408467, 0.61519726, 1.0, 1.8808158, 11.862615, 29.447498, 84.016846, 164.79132};
    return sec_to_years(planet_years[planet], seconds);
}