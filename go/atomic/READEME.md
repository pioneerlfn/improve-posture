# 为何这里不需要atomic

## 0x00

今天看了这篇[Go 标准库源码学习（一）详解短小精悍的 Once](https://mp.weixin.qq.com/s/Lsm-BMdKCKNQjRndNCLwLw) 之后，感觉对Go语言中的`atomic`理解又进一步，记录一下。


## 0x01 sync.Once

上面这篇文章集中讨论了`sync.Once`的实现，咱们直奔主题,先看下完整代码:

```Go

type Once struct {
	done uint32
	m    Mutex
}

func (o *Once) Do(f func()) {
    // 1
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
    // 2    
    o.m.Lock()
    defer o.m.Unlock()
    // 3
	if o.done == 0 {
        // 4
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

```

主要关心`注释1,3,4`三个地方对`done`读写的时候，为啥`1和4`需要用`atomic`， 而`3`这个地方不需要呢？

## 0x02 内存模型

主要是由Go的[内存模型](https://golang.org/ref/mem) 决定的.Go中不同goroutine之间，读写语义的`可观测性`需要一定的同步手段来保证。官博列出以下几种:

- Initialization
- Goroutine creation
- Goroutine destruction
- Channel communication
- Locks
- Once

没有提到`atomic`。但其实`atomic`也是能保证`可观测性的`,而且上面这几种手段也往往需要依赖`atomic`. `atomic`的代码，编译后的cpu指令具有`LOCK`前缀(x86), 说明是锁总线的，是Go里[最强的内存模型](https://github.com/cch123/golang-notes/blob/master/memory_barrier.md)


## 0x03 Lock与Unlock的happens-before

根据Go语言规范, `Unlock()`返回`happens-before` `Lock()`返回.
我们可以设想有`A`和`B`两个gorotine同时执行`Do`, 在`位置1`处检测的时候条件都为真，二者都进入`doSlow`.不过两者总有一个先那到锁，我门假设是`A`， 那么`B`此时便阻塞在`位置2`(`Lock`)这个地方了.

等`A`执行完`f`之后，调用`defer Unlock`, 因为`Unlock happens-before Lock`, 所以`位置4`这个地方的写操作，当`goroutine B` 从`Lock`返回，执行到`位置3`的时候y一定可以观测到`goroutine A`之前在`位置4`的写操作，即使`位置4`不是`atomic`也可以。

可见，由于`Lock`提供的`happens-before`语义保证，`位置3`这个地方的读操作是不需要用到`atomic`。


## 0x04 atomic

那`位置1`和`位置4`这两个个地方的`atomic`存在理由又是什么呢？

考虑这样一种情况, `goroutine A`刚执行完`位置4`处的写操作，`goroutine B`还被阻塞在`Lock`调用上，这时又有`goroutine C`再次进入`位置1`。

 假设`位置1`和`位置4`没有用`atomic`,而是普通读写。那么因为不同goroutine间对同一变量不存在同步手段，Go语言并不保证`goroutine A`的写操作能被`goroutine C`读到，导致`goroutine C`继续进入慢路径，这不符合我们的要求。

 所以，`位置1`和`位置4`的`atomic`是为了解决`gorouine A`和`goroutine C`这种情况下的可观测性的。

## 0x05 小结

本文详细考察了`sync.Once`的实现中，几处读写对于是否使用`atomic`时的考虑，加深了Go语言内存模型的理解。


## 0x06 参考资料

- [The Go Memory Model](https://golang.org/ref/mem#tmp_2) (官方)
- [Go 标准库源码学习（一）详解短小精悍的 Once](https://mp.weixin.qq.com/s/Lsm-BMdKCKNQjRndNCLwLw)
- [memory barrier](https://github.com/cch123/golang-notes/blob/master/memory_barrier.md) (曹春晖老师)
- [Cache coherency primer](https://fgiesen.wordpress.com/2014/07/07/cache-coherency/)

> 本文完。
