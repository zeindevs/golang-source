//go:generate g++ -c src/mylib.cpp -o lib/mylib.o
//go:generate ar rcs lib/libmylib.a lib/mylib.o
//go:generate g++ -c src/raywrap.cpp -o lib/raywrap.o -I./include -L./lib -lraylib
//go:generate ar rcs lib/libraywrap.a lib/raywrap.o
package main

/*
#cgo CXXFLAGS: -I./src -I./include
#cgo LDFLAGS: -L./lib -lmylib -lraywrap -lraylib -lm
#include <stdlib.h>
#include "src/mylib.hpp"
#include "src/raywrap.hpp"
*/
import "C"
import "unsafe"

func main() {
	println("Add:", C.add(10, 20))
	title := C.CString("Hello from Go + C++ + Raylib")
	defer C.free(unsafe.Pointer(title))

	C.rl_init_window(800, 450, title)

	for !bool(C.rl_should_close()) {
		C.rl_begin()
		C.rl_clear(245, 245, 245, 255)
		C.rl_draw_text(C.CString("Hello Window!"), 200, 200, 24)
		C.rl_end()
	}

	C.rl_close()
}
