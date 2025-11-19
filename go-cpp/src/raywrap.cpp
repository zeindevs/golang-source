#include <raylib.h>
#include "raywrap.hpp"

void rl_init_window(int width, int height, const char* title) {
    InitWindow(width, height, title);
}
void rl_begin() {
    BeginDrawing();
}
void rl_clear(int r, int g, int b, int a) {
    ClearBackground((Color){(unsigned char)r,(unsigned char)g,(unsigned char)b,(unsigned char)a});
}
void rl_draw_text(const char* txt, int x, int y, int size) {
    DrawText(txt, x, y, size, BLACK);
}
void rl_end() {
    EndDrawing();
}
bool rl_should_close() {
    return WindowShouldClose();
}
void rl_close() {
    CloseWindow();
}

