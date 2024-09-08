package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arrPtr := unsafe.Pointer(&arr[0])
	for i := 0; i < len(arr); i++ {
		next := (*int)(unsafe.Pointer(uintptr(arrPtr) + uintptr(i)*unsafe.Sizeof(arr[0])))
		fmt.Println(*next)
	}
}

func sample() {
	i := 10
	iptr := unsafe.Pointer(&i)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(*(*int)(iptr))
	// fmt.Println(*(*string)(iptr))
}
