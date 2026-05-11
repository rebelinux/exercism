#include "space_age.h"
#include "stdio.h"
#include "math.h"

double sec_to_years(double planetinyears, int64_t secs) {
    return secs / ((double)round(365.25 * planetinyears) * 86400);
}

float age(planet_t planet, int64_t seconds) {
    double year = 0.0;
    switch (planet)
    {
        case EARTH:
            year = seconds / 31557600;
            break;
        case MERCURY:{
            year = sec_to_years(0.2408467, seconds);
            break;
        }
        case VENUS:{
            year = sec_to_years(0.61519726, seconds);
            break;
        }
        case JUPITER:{
            year = sec_to_years(11.862615, seconds);
            break;
        }
        case MARS:{
            year = sec_to_years(1.8808158, seconds);
            break;
        }
        case SATURN:{
            year = sec_to_years(29.447498, seconds);
            break;
        }
        case URANUS:{
            year = sec_to_years(84.016846, seconds);
            break;
        }
        case NEPTUNE:{
            year = sec_to_years(164.79132, seconds);
            break;
        }
        default:
            year = -1;
            break;
    }
    return year;
}


// void main() {
//     float year = age(MERCURY, 2134835688);
//     printf("%.2f\n", year);
// }