#include "rna_transcription.h"
#include <stdlib.h>
#include "string.h"

char *to_rna(const char *dna) {
    char* rna = malloc(sizeof(dna));
    size_t i;
    for (i = 0; i < strlen(dna); i++) {
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
 