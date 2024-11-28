#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <dirent.h>
#include <ctype.h>

void rot13(char *input, char *output) {
    while (*input) {
        if (isalpha(*input)) {
            char offset = islower(*input) ? 'a' : 'A';
            *output = ((*input - offset + 13) % 26) + offset;
        } else {
            *output = *input;
        }
        input++;
        output++;
    }
    *output = '\0';
}

void process_file(const char *file_path) {
    FILE *file = fopen(file_path, "r");
    if (!file) {
        perror("Erro ao abrir o arquivo");
        return;
    }

    char **lines = NULL;
    size_t line_count = 0;

    char line[1024];
    while (fgets(line, sizeof(line), file)) {
        line[strcspn(line, "\n")] = '\0';
        char *obfuscated = malloc(strlen(line) + 1);
        if (!obfuscated) {
            perror("Erro ao alocar memória");
            fclose(file);
            return;
        }
        rot13(line, obfuscated);
        char **temp = realloc(lines, sizeof(char *) * (line_count + 1));
        if (!temp) {
            perror("Erro ao redimensionar memória");
            fclose(file);
            return;
        }
        lines = temp;
        lines[line_count++] = obfuscated;
    }

    fclose(file);

    file = fopen(file_path, "w");
    if (!file) {
        perror("Erro ao reabrir o arquivo para escrita");
        for (size_t i = 0; i < line_count; i++) {
            free(lines[i]);
        }
        free(lines);
        return;
    }

    for (size_t i = 0; i < line_count; i++) {
        fprintf(file, "%s\n", lines[i]);
        free(lines[i]);
    }

    free(lines);
    fclose(file);

    printf("Arquivo modificado: %s\n", file_path);
}

int main(int argc, char *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "Uso: %s <caminho_do_diretorio>\n", argv[0]);
        return 1;
    }

    const char *directory_path = argv[1];

    DIR *dir = opendir(directory_path);
    if (!dir) {
        perror("Erro ao abrir o diretório");
        return 1;
    }

    struct dirent *entry;
    char file_path[1024];

    while ((entry = readdir(dir))) {
        if (entry->d_type == DT_REG) {
            snprintf(file_path, sizeof(file_path), "%s/%s", directory_path, entry->d_name);
            process_file(file_path);
        }
    }

    closedir(dir);
    return 0;
}

