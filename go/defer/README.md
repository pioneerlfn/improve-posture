
## defer与return的执行顺序

`return`不是原子语句，`return expr`可以分解为:

1. 返回值=expr
2. return

如果在函数中还有`defer`，那么执行顺序是:
1. 返回值=expr
2. 被defer的函数(返回值有可能在defer中被修改)
3. return

看个例子🌰：

```go

package main

import "fmt"

var gl string

func main() {	
	gl = "main"
	fmt.Println(gl)
	foo()
	fmt.Println(gl)
}

func foo() error {	
	gl = "foo"
	defer df()
	fmt.Println(gl)
	return bar()
}

func bar() error {	
	gl = "bar"
	fmt.Println(gl)
	return nil
}

func df() {	
	gl = "defer"
}

```
```bash
➜ ~  go run defer.go 
main
foo
bar
defer
```


## 推荐阅读
- [Golang中的Defer必掌握的7知识点](https://zhuanlan.zhihu.com/p/115472856)