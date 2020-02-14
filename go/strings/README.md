# 如何多块好省拼接字符串

拼接字符串这个问题，`飞雪无情老师`写过三篇博客阐述过:

- [Go语言字符串高效拼接（一）](https://www.flysnow.org/2018/10/28/golang-concat-strings-performance-analysis.html)
- [Go语言字符串高效拼接（二）](https://www.flysnow.org/2018/11/05/golang-concat-strings-performance-analysis.html)
- [Go语言字符串高效拼接（三）](https://www.flysnow.org/2018/11/11/golang-concat-strings-performance-analysis.html)

可以知道，最长用的字符串拼接方式有：

- `+`
- `strings.Join`
- `fmt.Sprint`
- `bytes.buffer`
- `strings.Builder`

通过压测可以知道：

- `+`连接适用于短小的、常量字符串（明确的，非变量），因为编译器会给我们优化
- `Join`是比较统一的拼接，不太灵活. 不过个人觉得在综合考虑速度和内存开辟，`Join`的表现是最好的。
- `fmt和buffer`基本上不推荐
- `builder`从性能和灵活性上，都是上佳的选择

实际上，`Join`和`Builder`表现都很不错。但这两个方法的使用侧重点有些不一样：

- 如果有现成的数组、切片那么可以直接使用`Join`, `Join`还是定位于有现成切片、数组的（毕竟拼接成数组也要时间），并且使用固定方式进行分解的，比如逗号、空格等，局限比较大
- 但是如果没有，并且追求灵活性拼接，还是选择Builder

第三篇博客讲了`Builder`的优化技巧，如果需要频繁写入大量的`string`,可能导致底层数组多次扩容，从而严重影响其性能。若我们能在事先大概知道总的内存需求，可以调用`Builder.Grow()`一次性开辟足够内存，大大提升其性能。
