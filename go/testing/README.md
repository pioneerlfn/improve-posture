# Go学习笔记——测试技巧备忘


## subtest

subtest主要解决`表驱动测试`方法中，无法对一个测试函数中的，待测试cases数组中特定case单独测试的问题。

通过给每一个测试case起一个名字，我们可以单独测试case数组中特定一个或一批case. 

比如，通过下列命令，可以单独测试`TestOlder`函数中名为`FirstOlderThanSecond`的case.

```go
$ go test -v -count=1 -run="TestOlder/FirstOlderThanSecond"

```
注意，我们在测试命令中用到了`-run=`开关，用来选择`*_test.go`文件中待测试的函数及测试case. `-run=`选择遵守`正则匹配`规则.

> Read More: [Testing in Go: Subtests](https://ieftimov.com/post/testing-in-go-subtests/)
