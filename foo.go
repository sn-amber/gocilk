package main

// #include "foo.h"
// #cgo CFLAGS: -fcilkplus
// #cgo LDFLAGS: -lcilkrts
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	arr := make([]int32, 10, 10)
	var i int32
	for i = 0; i < 10; i++ {
		arr[i] = i
	}

	fmt.Printf("array: %v\n", arr)

	arr0 := unsafe.Pointer(&arr[0])
	s := C.sum((*C.int)(arr0), 0, 10)

	fmt.Printf("sum: %d\n", s)
}
