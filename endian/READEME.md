# 理解机器大小端

本文记录一下对机器大小端的测试，并借此机会熟悉一下Go中指针的使用。代码如下:

```Go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 0x12345678
	size := unsafe.Sizeof(a)
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("size of a: %d\n", size)

	var p = uintptr(unsafe.Pointer(&a))
	var p1 *byte
	for i := 0; i < int(size); i++ {
		p1 = (*byte)(unsafe.Pointer(p))
		fmt.Printf("%p: 0x%x\n", p1, *p1)
		p = p + uintptr(1)
	}
}

```

运行程序，结果如下:
```
Type of a: int
size of a: 8
0xc000076f08: 0x78
0xc000076f09: 0x56
0xc000076f0a: 0x34
0xc000076f0b: 0x12
0xc000076f0c: 0x0
0xc000076f0d: 0x0
0xc000076f0e: 0x0
0xc000076f0f: 0x0

```

可见`int`类型的`a = 0x12345678`在内存中的布局如下图所示:
![endian](./endian.png)

我的机器是`x86`架构，`x86`是小端架构。这样我们就知道了，小端架构的机器，多余多字节数据类型，最重要的位存于高内存地址，最不重要的存于低地址内存。
