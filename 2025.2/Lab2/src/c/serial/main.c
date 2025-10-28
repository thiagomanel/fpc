#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>
#include <sys/time.h>

#define MAX_STUDENTS 10000
#define MAX_CLASSES 100

typedef struct {
    char student_id[16];
    int class_id;
    double final_grade;
    int has_grade;
} Student;

typedef struct {
    int num_classes;
    int num_students_per_class;
    int total_students;
    Student registry[MAX_STUDENTS];
    char *professors[2];
    int class_ids[MAX_CLASSES];
} Classes;

/* Generate random student ID */
void generate_student_id(char *buffer) {
    static int next_id = 1;  
    sprintf(buffer, "%d", next_id);  
    next_id++;                       
}

/* Generate random grade */
double generate_student_grade() {
    return ((double)rand() / RAND_MAX) * 10.0;
}

/* Sleep for random float seconds between min and max */
void random_sleep(double min_sec, double max_sec) {
    double duration = min_sec + ((double)rand() / RAND_MAX) * (max_sec - min_sec);
    usleep((int)(duration * 1000000)); // convert seconds to microseconds
}

/* Initialize semester registry */
void init_registry(Classes *c) {
    int index = 0;
    for (int class_id = 1; class_id <= c->num_classes; class_id++) {
        for (int i = 0; i < c->num_students_per_class; i++) {
            generate_student_id(c->registry[index].student_id);
            c->registry[index].class_id = class_id;
            c->registry[index].final_grade = 0.0;
            c->registry[index].has_grade = 0;
            index++;
        }
    }
    c->total_students = index;
}

int get_students_in_class(Classes *c, int class_id, Student *output[]) {
    int count = 0;
    for (int i = 0; i < c->total_students; i++) {
        if (c->registry[i].class_id == class_id) {
            output[count++] = &c->registry[i];
        }
    }
    return count;
}

/* Process grades for one class */
void process_grades(Classes *c, int class_id) {
    const char *professor = c->professors[(class_id - 1) % 2];
    printf("\n%s's class (%d) — starting grading...\n", professor, class_id);
    fflush(stdout);
    random_sleep(0.5, 1.0);

    Student *students[MAX_STUDENTS];
    int count = get_students_in_class(c, class_id, students);

    for (int i = 0; i < count; i++) {
        Student *s = students[i];
        fflush(stdout);
        random_sleep(0.2, 0.5);

        double grade = generate_student_grade();
        s->final_grade = grade;
        s->has_grade = 1;

        printf("%s corrected Student %s from class %d - Grade: %.2f\n", professor, s->student_id, class_id, grade);
        fflush(stdout);
        random_sleep(0.1, 0.3);
    }

    printf("%s's class %d grades successfully processed!\n\n", professor, class_id);
}

/* Print registry per class */
void registry_to_string(Classes *c, int class_id) {
    const char *professor = c->professors[(class_id - 1) % 2];
    Student *students[MAX_STUDENTS];
    int count = get_students_in_class(c, class_id, students);

    printf("\n*********** %s's Class %d***********\n", professor, class_id);
    for (int i = 0; i < count; i++) {
        Student *s = students[i];
        printf("student_id: %s, class_id: %d, final_grade: ", s->student_id, s->class_id);
        if (s->has_grade)
            printf("%.2f\n", s->final_grade);
        else
            printf("None\n");
    }
}

/* Initialize Classes structure */
void init_classes(Classes *c, int num_classes, int num_students_per_class) {
    c->num_classes = num_classes;
    c->num_students_per_class = num_students_per_class;
    c->professors[0] = "Prof1";
    c->professors[1] = "Prof2";
    for (int i = 0; i < num_classes; i++) {
        c->class_ids[i] = i + 1;
    }
    init_registry(c);
}

/* --------------------------- MAIN --------------------------- */
int main(int argc, char *argv[]) {
    srand((unsigned int)time(NULL));

    if (argc < 3) {
        printf("Usage: %s <num_classes> <num_students_per_class>\n", argv[0]);
        return 1;
    }

    int num_classes = atoi(argv[1]);
    int num_students_per_class = atoi(argv[2]);

    if (num_classes > MAX_CLASSES || num_classes <= 0 ||
        num_students_per_class <= 0 || num_students_per_class * num_classes > MAX_STUDENTS) {
        printf("Invalid input values.\n");
        return 1;
    }

    Classes semester;
    init_classes(&semester, num_classes, num_students_per_class);

    printf("======================= Welcome to 2025.2's UFCG Semester =======================\n");

    for (int i = 0; i < semester.num_classes; i++) {
        process_grades(&semester, semester.class_ids[i]);
    }

    for (int i = 0; i < semester.num_classes; i++) {
        registry_to_string(&semester, semester.class_ids[i]);
    }

    printf("======================= The End =======================\n");

    return 0;
}
