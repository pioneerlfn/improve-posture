package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 []int
	var s2 = []int{}

	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // false

	p1 := *(*[3]int)(unsafe.Pointer(&s1))
	p2 := *(*[3]int)(unsafe.Pointer(&s2))

	fmt.Println(p1) // [0 0 0]
	fmt.Println(p2) // [824634150592 0 0]
}
