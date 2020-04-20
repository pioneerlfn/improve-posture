## Strings.EqualFold

比较字符串可以用`==`, `Strings.Compare`, `Strings.EqualFold`.

前两种方法很常见，第三种`EqualFold`最近才新注意到，可以理解为`case-insensive`的比较。

例子🌰:

```Go
// filename: equalfold.go
package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println("Go" == "go")
	fmt.Println(strings.Compare("Go", "go"))
	fmt.Println(strings.EqualFold("Go", "go"))	
}
```

```bash
go run equalfold.go
false
-1
true
```

## 推荐阅读:
[strings: EqualFold doc could be clearer #33447](https://github.com/golang/go/issues/33447)
