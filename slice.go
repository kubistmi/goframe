package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main(){
	a := []int{1,2,3,4}

	b := make([]int, 0, 4)
	b = append(b, a[0])
	b = append(b, a[2])

	shA := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	shB := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	fmt.Println(shA)
	fmt.Println(shB)
}