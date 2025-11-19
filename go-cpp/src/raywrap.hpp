#pragma once

#ifdef __cplusplus
extern "C" {
#endif

void rl_init_window(int width, int height, const char* title);
void rl_begin();
void rl_clear(int r, int g, int b, int a);
void rl_draw_text(const char* txt, int x, int y, int size);
void rl_end();
bool rl_should_close();
void rl_close();

#ifdef __cplusplus
}
#endif
