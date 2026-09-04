#include "two_fer.h"


void two_fer(char *buffer, const char *name) {
    if (name == NULL)
    {
        snprintf(buffer, BUFFER_SIZE, "One for you, one for me.");
    } else {
        snprintf(buffer, BUFFER_SIZE, "One for %s, one for me.", name);
    } 
}
