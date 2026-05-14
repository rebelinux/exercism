#include "allergies.h"

bool is_allergic_to(allergen_t allergen, int count) {
    return (count >> allergen) & 1;
}

allergen_list_t get_allergens(int count) {
        allergen_list_t t = {0};
        for (int i = 0; i < ALLERGEN_COUNT; i++) {
            bool allergic = is_allergic_to(i, count);
            if (allergic) {
                t.allergens[i] = allergic;
                ++t.count;
            }
        }
        return t;
}
