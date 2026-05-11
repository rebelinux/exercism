#include "grade_school.h"
#include <string.h>
#include "stdlib.h"


roster_t init_roster(roster_t *roster) {
    *roster = (roster_t){0};
    return *roster;
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

int compareByGrade(const void *a, const void *b) {
    // Cast void pointers to struct Student pointers
    student_t *studentA = (student_t *)a;
    student_t *studentB = (student_t *)b;
    
    // Return (A - B) for ascending order
    return (studentA->grade - studentB->grade);
}

int compareByName(const void *a, const void *b) {
    // Cast void pointers to struct Student pointers
    student_t *studentA = (student_t *)a;
    student_t *studentB = (student_t *)b;
    
    // Return (A - B) for ascending order
    return strcmp(studentA->name, studentB->name);
}

bool add_student(roster_t *roster, char *name, int grade) {
    bool result;
    if (roster->count < 1) {
        roster->students[roster->count].grade = grade;
        strcpy(roster->students[0].name, name);
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
    qsort(roster->students, roster->count, sizeof(student_t), compareByName);
    qsort(roster->students, roster->count, sizeof(student_t), compareByGrade);
    return result;
}

roster_t get_grade(roster_t *roster, uint8_t grade) {
    roster_t r;
    init_roster(&r);
    for (size_t i = 0; i < roster->count; i++) {
        if (roster->students[i].grade == grade) {
            r.students[r.count].grade = roster->students[i].grade;
            strcpy(r.students[r.count].name,roster->students[i].name);
            r.count++;
        }
    }
    return r;
}

// int main() {
//     roster_t actual;
//     init_roster(&actual);
//     add_student(&actual, "Peter", 3);
//     add_student(&actual, "Zoe", 2);
//     add_student(&actual, "Alex", 1);
//     add_student(&actual, "Paul", 1);
//     roster_t grade1 = get_grade(&actual, 1);
//     for (int i = 0; i < actual.count; i++) {
//         printf("All Student");
//         printf("Name: %s, Grade: %d\n", actual.students[i].name, actual.students[i].grade);
//     }
//     for (int i = 0; i < grade1.count; i++) {
//         printf("Grade 1 Student");
//         printf("Name: %s, Grade: %d\n", grade1.students[i].name, grade1.students[i].grade);
//     }
// }  