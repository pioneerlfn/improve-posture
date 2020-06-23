## 关于defer你应该知道的

### defer与return的执行顺序

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

### defer函数参数evaluate时机
The arguments of a deferred function call or a goroutine function call `are all evaluated at the moment when the function call is invoked.`
- For a deferred function call, the invocation moment is `the moment when it is pushed into the defer-call stack of its caller goroutine.`
- For a goroutine function call, the invocation moment is the moment when the corresponding `goroutine is created.`

### 练习题
读者诸君可以试一下下面几道面试题:

```Go
package main

import "fmt"

func deferFunc1(i int) (t int) {
	t = i // i=1, t=1
	defer func() {
		t += 3
	}()
	return t
} // 4

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
} // 1

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
} // 3

func DeferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i) // 0
		fmt.Println(t) // 2
	}(t) // 0
	t = 1
	return 2 // t=2
}

func main() {
	fmt.Println(deferFunc1(1)) // 4
	fmt.Println(DeferFunc2(1)) // 1
	fmt.Println(DeferFunc3(1)) // 3
	DeferFunc4()               // 0, 2
}

```



## 推荐阅读
- [Golang中的Defer必掌握的7知识点](https://zhuanlan.zhihu.com/p/115472856)