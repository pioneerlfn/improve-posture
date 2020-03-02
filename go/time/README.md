# 谈一谈Go中的时区问题

## 几个概念

### Unix时间戳

> The unix time stamp is a way to track time as a running total of seconds. This count starts at the Unix Epoch on January 1st, 1970 at UTC. Therefore, the unix time stamp is merely **the number of seconds between a particular date and the Unix Epoch. It should also be pointed out that this point in time technically does not change no matter where you are located on the globe.**

可以看出`unix 时间戳`是一个绝对值，与使用者所在的时区没有关系。如图1所示:
![unix timestamp](timestamp.png)

### Wall Clock vs Monotonic Clock

> Operating systems provide both a “wall clock,” which is subject to changes for clock synchronization, and a “monotonic clock,” which is not. The general rule is that the wall clock is for telling time and the monotonic clock is for measuring time. Rather than split the API, in this package the Time returned by time.Now contains both a wall clock reading and a monotonic clock reading; later time-telling operations use the wall clock reading, but later time-measuring operations, specifically comparisons and subtractions, use the monotonic clock reading.

- Wall clock(time)就是我们一般意义上的时间，就像墙上钟所指示的时间。

- Monotonic clock(time)字面意思是单调时间，实际上它指的是从某个点开始后（比如系统启动以后）流逝的时间，jiffies一定是单调递增的！

而特别要强调的是计算两个时间点的差值一定要用Monotonic clock(time)，因为Wall clock(time)是可以被修改的，比如计算机时间被回拨（比如校准或者人工回拨等情况），或者闰秒（ leap second），会导致两个wall clock(time)可能出现负数。（因为操作系统或者上层应用不一定完全支持闰秒，出现闰秒后系统时间会在后续某个点会调整为正确值，就有可能出现时钟回拨（当然也不是一定，比如ntpdate就有可能出现时钟回拨，但是ntpd就不会））

> Read More:
- [Clock and Monotonic Clock](https://taobig.org/?p=752)
- [time document(Go)](https://golang.org/pkg/time/)


## Go中time包中的重要类型

### Time类型

[time包](https://golang.org/pkg/time/) 中最重要的类型是`Time`, 定义如下：
```go
type Time struct {
	// wall and ext encode the wall time seconds, wall time nanoseconds,
	// and optional monotonic clock reading in nanoseconds.
	//
	// From high to low bit position, wall encodes a 1-bit flag (hasMonotonic),
	// a 33-bit seconds field, and a 30-bit wall time nanoseconds field.
	// The nanoseconds field is in the range [0, 999999999].
	// If the hasMonotonic bit is 0, then the 33-bit field must be zero
	// and the full signed 64-bit wall seconds since Jan 1 year 1 is stored in ext.
	// If the hasMonotonic bit is 1, then the 33-bit field holds a 33-bit
	// unsigned wall seconds since Jan 1 year 1885, and ext holds a
	// signed 64-bit monotonic clock reading, nanoseconds since process start.
	wall uint64
	ext  int64

	// loc specifies the Location that should be used to
	// determine the minute, hour, month, day, and year
	// that correspond to this Time.
	// The nil location means UTC.
	// All UTC times are represented with loc==nil, never loc==&utcLoc.
	loc *Location
}

```
总体来讲，`Time`结构作为time包中的核心结构体，主要包括两部分：

1. 用来编码clock的`wall`和`ext`字段；
2. 用来表征时区的`loc`字段


根据注释我们可以知道：
- `wall`和`ext`一起编码`wall clock`和`monotonic clock`;

- `monototic clock`是可选的，且`wall`字段的最高位用来表征是否包含`monotonic clock`, *MSB = 1*代表包含，*MSB = 0*代表不包含;

- 如果`wall`的最高位为1(*MSB=1*)，那么`wall`从最高位以下的33位以无符号数的形式，代表从`1985年1月1日零时`开始的秒数(注意:因为不是从1970年开始，所以不是Unix时间戳), 最后30位代表纳秒数;`ext`解释为位从进程启动以来的纳秒数，是为`monotonic clock`;

- 如果`wall`最高位为0(*MSB = 0*), 那么往下的33位无效置零，剩余30位仍然表示纳秒值;而`ext`则表示从公元1年1月1日零时开始的秒数(注意:也不是unix时间戳)

> 可见，在`Time`中没有直接存储`unix时间戳`.

### 时区

`Time`类型中的第三个字段`loc`字段用来表征使用的时区。

根据前面的分析我们可以知道，时间是由`wall`和`ext`通过前述规则表示的，本身是个绝对值。但是对同一个时间的表示(format),可以有很多形式。

比如`time.Now()`代表当前时间，虽然用格林尼治时区或是北京时区或是美国纽约时区表示出来不一样，但在时间长河中，确实是同一个点。看下面的例子：

```go

func main() {
	now := time.Now() // #1
	fmt.Println(now) // #2 2020-03-03 00:15:54.801931 +0800 CST m=+0.000092900
	fmt.Println(now.UTC()) // #3 2020-03-02 16:15:54.801931 +0000 UTC

	// unix timestamp
	fmt.Println(now.Unix()) // #4 1583165754
	fmt.Println(now.UTC().Unix()) // #5 1583165754
}

```

这里有个问题不知道你注意到没有：在上面的代码中，时区信息究竟是怎么注入的呢？什么时候注入的？下面我们回答一下这个问题。

看下`time.Now`的代码:

```Go
// Now returns the current local time.
func Now() Time {
	sec, nsec, mono := now()
	mono -= startNano
	sec += unixToInternal - minWall
	if uint64(sec)>>33 != 0 {
		return Time{uint64(nsec), sec + minWall, Local} // 时区在这里,Local是一个包级别全局变量, *Location类型
	}
	return Time{hasMonotonic | uint64(sec)<<nsecShift | uint64(nsec), mono, Local}
}
 
```
`Local`实际上是一个指向包级别非导出变量`localLoc`

```Go
var Local *Location = &localLoc
// localLoc is separate so that initLocal can initialize
// it even if a client has changed Local.
var localLoc Location

```
我们又看到, `localLoc`没有被显示初始化的，也就是说刚开始是`Location`的零值，没有本地时区信息。

也就是说
```Go
time.Now()
```
这行代码拿到的`Time`仍然不包含本地时区信息。

那么为啥
```Go
fmt.Println(time.Now())
```
打印出来的是本地时区的时间表示(format)呢?

通过研究，我们发现,`initLocal()`函数负责初始化时区信息:

```Go
func initLocal() {
	// consult $TZ to find the time zone to use.
	// no $TZ means use the system default /etc/localtime.
	// $TZ="" means use UTC.
	// $TZ="foo" means use /usr/share/zoneinfo/foo.

	tz, ok := syscall.Getenv("TZ")
	switch {
	case !ok:
		z, err := loadLocation("localtime", []string{"/etc/"})
		if err == nil {
			localLoc = *z
			localLoc.name = "Local"
			return
		}
	case tz != "" && tz != "UTC":
		if z, err := loadLocation(tz, zoneSources); err == nil {
			localLoc = *z
			return
		}
	}

	// Fall back to UTC.
	localLoc.name = "UTC"
}

```
看一下上面代码就能知道加载本地时区的逻辑。
`initLocal()`负责初始化`localLoc`, 这里用到了`单例模式`:

```Go
// localLoc is separate so that initLocal can initialize
// it even if a client has changed Local.
var localLoc Location
var localOnce sync.Once

func (l *Location) get() *Location {
	if l == nil {
		return &utcLoc
	}
	if l == &localLoc {
		localOnce.Do(initLocal)
	}
	return l
}

```

所以，我们就能知道当我们打印的时候会调用`Time.String()`, `String`最后会调用到`Location.get()`，从而实现本地时区信息的`lazy load`. 

### 小结

本文主要目的在于分析清楚`Time`类型中，本地时区信息的加载时机。使用`time`包的注意事项，下面几篇文章写得很好，请仔细阅读:

- [深入理解GO时间处理(time.Time)](https://www.imhanjm.com/2017/10/29/%E6%B7%B1%E5%85%A5%E7%90%86%E8%A7%A3golang%E6%97%B6%E9%97%B4%E5%A4%84%E7%90%86(time.time)/)

- [golang time 包的坑](https://blog.wolfogre.com/posts/trap-of-golang-time/)



