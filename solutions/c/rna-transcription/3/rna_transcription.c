#include "rna_transcription.h"
#include <stdlib.h>
#include "string.h"

char *to_rna(const char *dna) {
    size_t slen = strlen(dna);
    // "sizeof(dna) + 1" display a segfault :(
    char *rna = malloc(slen * sizeof(char) + 1);
    size_t i;
    for (i = 0; i < slen; i++) {
        switch (dna[i])
        {
            case 'G':
                rna[i] = 'C';
                break;
            case 'C':
                rna[i] = 'G';
                break;
            case 'T':
                rna[i] = 'A';
                break;
            case 'A':
                rna[i] = 'U';
                break;
            default:
                rna[i] = ' ';
                break;
        }
    }
    rna[i] = '\0';
    return rna;
}