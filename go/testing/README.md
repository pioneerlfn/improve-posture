# Go学习笔记——测试技巧备忘

## table driven test

Go中测试一般使用`表驱动测试`.

> Read More:

- [Testing in Go: Table-Driven Tests](https://ieftimov.com/post/testing-in-go-table-driven-tests/)
- [Dava Cheney: Prefer table driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

## subtest

subtest主要解决`表驱动测试`方法中，无法对一个测试函数中的，待测试cases数组中特定case单独测试的问题。

通过给每一个测试case起一个名字，，利用`testing`包中的`t.Run`, 我们可以单独测试case数组中特定一个或一批case.
使用`t.Run`还有一个好处，就是我们可以再每个测试case开始前和结束后插入`setup`和`teardown`函数，进行一些状态设置和清除工作.

比如，通过下列命令，可以单独测试`TestOlder`函数中名为`FirstOlderThanSecond`的case.

```go
$ go test -v -count=1 -run="TestOlder/FirstOlderThanSecond"
...
```

注意，我们在测试命令中用到了`-run=`开关，用来选择`*_test.go`文件中待测试的函数及测试case. `-run=`选择遵守`正则匹配`规则.

> Read More: [Testing in Go: Subtests](https://ieftimov.com/post/testing-in-go-subtests/)

## TestMain

如果要在测试文件的所有测试函数被测试之前与之后执行一些`setup`以及`teardown`操作，可以使用`TestMain`,一次典型的`TestMain`过程如下:

```go
func TestMain(m *testing.M) {
    setup()
    code := m.Run() 
    shutdown()
    os.Exit(code)
}
```

> Read More: [TestMain—What is it Good For?](http://cs-guy.com/blog/2015/01/test-main/)

## 测试main包中函数

对于测试包go test是一个的有用的工具，但是稍加努力我们也可以用它来测试可执行程序。如果一个包的名字是 main，那么在构建时会生成一个可执行程序，不过main包可以作为一个包被测试器代码导入.

比如我们想要测试`main`包有一个函数`echo`。虽然是main包，也有对应的main入口函数，但是在测试的时候main包只是`TestEcho`测试函数导入的一个普通包，里面main函数并没有被导出，而是被忽略的。

## fixture



## Dependency Injection(依赖注入)

主要用到两种方式，一种是函数闭包，另一种是接口。
接口的方法更推荐一点。

## test doubles

之前没听过这个词，以为`mock`就是全部。在看了[Testing in Go: Test Doubles by Example](https://ieftimov.com/post/testing-in-go-test-doubles-by-example/)这篇文章才对这个概念有了基本的了解。 `test doubles`的含义比较广，包括`Dummies, mocks, stubs, fakes, and spies`. 这些方法虽然看起来很像，但是又有一些微妙的不同。也算是各种不同的`pattern`. 能清晰地分辨不同的pattern, 窃以为是码农很重要的能力。

> Read More: [Testing in Go: Test Doubles by Example](https://ieftimov.com/post/testing-in-go-test-doubles-by-example/)

## golden files

当输出的内容太多太复杂时，hardcode的办法就有些力不从心了。这个时候就会就可以用`golden files`的方法，将test的输出内容保存在golen文件中.

## 参考文章

- `go help test`
- [The Go Blog: Using Subtests and Sub-benchmarks](https://blog.golang.org/subtests)
- [The Go Programming Language: Ch11-testing](https://yar999.gitbooks.io/gopl-zh/content/ch11/ch11.html)
- [Testing in Go: Table-Driven Tests](https://ieftimov.com/post/testing-in-go-table-driven-tests/)
- [Dava Cheney: Prefer table driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
