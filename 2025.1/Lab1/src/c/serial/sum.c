#include <stdio.h>
#include <stdlib.h>

int do_sum(const char *path) {
    FILE *file = fopen(path, "rb");  // open in binary mode
    if (file == NULL) {
        return -1;  // indicate error
    }

    int sum = 0;
    int byte;

    while ((byte = fgetc(file)) != EOF) {
        sum += byte;
    }

    fclose(file);
    return sum;
}

int main(int argc, char *argv[]) {
    for (int i = 1; i < argc; i++) {
        const char *path = argv[i];
        int sum = do_sum(path);

        if (sum >= 0) {
            printf("%s : %d\n", path, sum);
        }
    }

    return 0;
}
