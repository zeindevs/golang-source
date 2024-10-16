#include "libmain.h"
#include <iostream>

int main() {
  StartChannel();

  while (true) {
    int value = SubscribeChannel();
    if (value == -1) {
      break;
    }
    std::cout << "Received from Go channel: " << value << std::endl;
  }
  return 0;
}
