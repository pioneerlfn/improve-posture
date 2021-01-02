# 分隔字符串

`func Split(s, sep string) []string`

- Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
- If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
- If sep is empty, Split splits after each UTF-8 sequence. If both s and sep are empty, Split returns an empty slice.
- It is equivalent to SplitN with a count of -1.

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

	fmt.Printf("q%\n", strings.Split("", "/")) // []string{""}, len = 1, cap = 1.
}
```

可见，对于字符串`“/home/lfn/”`， 分隔之后的切片会有4个元素，第一个和第四个都是`""`.
