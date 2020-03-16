# 分隔字符串

看下面例子:
```Go

package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "/home/lfn/"
	b := "/home/lfn"
	as := strings.Split(a, "/")
	bs := strings.Split(b, "/")
	fmt.Println("len(as):", len(as)) // 4
	fmt.Println("len(bs):", len(bs)) // 3
	fmt.Printf("as: %#v\n", as) //as: []string{"", "home", "lfn", ""}
	fmt.Printf("bs: %#v\n", bs) // bs: []string{"", "home", "lfn"}
}
```

可见，对于字符串`“/home/lfn/”`， 分隔之后的切片会有4个元素，第一个和第四个都是`""`.
