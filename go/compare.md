## 函数类型是不可比较的

1. 函数值是不可比较的，比较会产生panic;
2. 函数值可以和没有类型的 `nil`比较;
3. 由于函数值不可比较，所以它们不能作为字典的key.

举个例子:
```Go
package main

import "fmt"

func main() {
	f := func() {}
	compare(T{1, f}, T{2, f})
	compare(T{f, 1}, T{f, 2})
	fmt.Println(f == nil)
}

type T [2]interface{}

func compare(a, b T) {
	defer func() {
		if recover() != nil {
			fmt.Println("panic")
		}
	}()
	fmt.Println(a == b)
}

```
输出是：
```shell
false
panic
false
```