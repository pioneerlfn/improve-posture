package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s1 []int     // nil切片
	var s2 = []int{} // 空切片

	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // false

	p1 := *(*[3]int)(unsafe.Pointer(&s1))
	p2 := *(*[3]int)(unsafe.Pointer(&s2))

	fmt.Println(p1) // [0 0 0]
	fmt.Println(p2) // [824634150592 0 0]
}

// 参考[深度解析 Go 语言中「切片」的三种特殊状态](https://juejin.cn/post/6844903712654098446)
