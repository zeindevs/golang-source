#include <stdio.h>
#include <windows.h>

typedef int (*AddFunc)(int, int);
typedef int (*SubtractFunc)(int, int);

int main() {
  HMODULE h = LoadLibrary("mylib.dll");
  if (!h) {
    printf("Failed to load the DLL\n");
    return 1;
  }

  AddFunc Add = (AddFunc)GetProcAddress(h, "Add");
  SubtractFunc Subtract = (SubtractFunc)GetProcAddress(h, "Subtract");

  if (!Add || !Subtract) {
    printf("Failed to locate functions\n");
    return 1;
  }

  printf("Add(5, 3) = %d\n", Add(5, 3));
  printf("Subtract(5, 3) = %d\n", Subtract(5, 3));

  FreeLibrary(h);
  return 0;
}
