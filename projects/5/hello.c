#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <emscripten/emscripten.h>

EMSCRIPTEN_KEEPALIVE
char* generate_greeting(const char* name, int age, int repeat) {
    const char* format = "Hello, %s. You are %d years old.\n";
    int single_len = snprintf(NULL, 0, format, name, age);
    int total_len = single_len * repeat + 1;

    char* result = (char*)malloc(total_len);
    if (!result) return NULL;

    char* ptr = result;
    for (int i = 0; i < repeat; ++i) {
        ptr += sprintf(ptr, format, name, age);
    }

    return result;
}

EMSCRIPTEN_KEEPALIVE
void free_buffer(char* ptr) {
    free(ptr);
}
