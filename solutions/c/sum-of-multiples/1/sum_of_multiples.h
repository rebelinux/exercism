#ifndef SUM_OF_MULTIPLES_H
#define SUM_OF_MULTIPLES_H

#include <stddef.h>
#include <stdbool.h>

unsigned int sum(const unsigned int *factors, const size_t number_of_factors,
                 const unsigned int limit);
bool validate(int *array, size_t size, int number);

#endif
