#include "grade_school.h"
#include <string.h>
#include "stdlib.h"

void init_roster(roster_t *roster) {
    *roster = (roster_t){0};
}

bool match(roster_t *roster, char *name) {
    bool result = false;
    for (size_t i = 0; i < roster->count; i++) {
        int compare = strcmp(roster->students[i].name, name);
        if (compare == 0) {
            result = true;
            break;
        }
    }
    return result;
}

int compareByGradeAndName(const void *a, const void *b) {
    student_t *studentA = (student_t *)a;
    student_t *studentB = (student_t *)b;
    
    if (studentA->grade != studentB->grade)
        return studentA->grade - studentB->grade;
    return strcmp(studentA->name, studentB->name);
}

bool add_student(roster_t *roster, char *name, uint8_t grade) {
    bool result;
    if (roster->count < 1) {
        roster->students[roster->count].grade = grade;
        strcpy(roster->students[roster->count].name, name);
        roster->count++;
        result = true;
    } else {
        if (!match(roster, name)) {
            roster->students[roster->count].grade = grade;
            strcpy(roster->students[roster->count].name, name);
            roster->count++;
            result = true;
        } else {
            result = false;
        }
    }
    qsort(roster->students, roster->count, sizeof(student_t), compareByGradeAndName);
    return result;
}

roster_t get_grade(roster_t *roster, uint8_t grade) {
    roster_t r;
    init_roster(&r);
    for (size_t i = 0; i < roster->count; i++) {
        if (roster->students[i].grade == grade) {
            r.students[r.count] = roster->students[i];
            r.count++;
        }
    }
    return r;
} 