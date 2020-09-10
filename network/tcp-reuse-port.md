## SO_REUSEPORT备忘

### 0x00 epoll的惊群问题
为了理解`SO_REUSEPORT`，有必要先了解一下`epoll`的`惊群问题`
- [Epoll is fundamentally broken 1/2](https://idea.popcount.org/2017-02-20-epoll-is-fundamentally-broken-12/)

- 再看一下cloudfare这篇[Why does one NGINX worker take all the load?](https://blog.cloudflare.com/the-sad-state-of-linux-socket-balancing/)

### 0x01 SO—REUSEPORT是什么

kernel 3.9为`socket`引入了`SO_REUSEPORT`选项，目的是为了在进程间共享端口,由kernel实现各个进程间的负载均衡：
- 我们应该先读一遍LWN上[The SO_REUSEPORT socket option](https://lwn.net/Articles/542629/)，这篇文章也可以看作是对这一特性的官方介绍。

- 掘金小册`《深入理解 TCP 协议：从原理到实战》`中的[一台主机上两个进程可以同时监听同一个端口吗](https://juejin.im/book/6844733788681928712/section/6844733788832923661)对这个问题有很全面的讲解，也值得一读。

- [man page](https://man7.org/linux/man-pages/man7/socket.7.html) 中的简单介绍:
> Permits multiple AF_INET or AF_INET6 sockets to be bound to an identical socket address.  This option must be set on each socket (including the first socket) prior to calling bind(2)on the socket.  To prevent port hijacking, all of the pro‐
cesses binding to the same address must have the same effec‐tive UID.  This option can be employed with both TCP and UDP sockets.

> For TCP sockets, this option allows accept(2) load distribution in a multi-threaded server to be improved by using a dis‐tinct listener socket for each thread.  This provides improved load distribution as compared to traditional techniques such using a single accept(2)ing thread that distributes connections, or having multiple threads that compete to accept(2) from the same socket.

> For UDP sockets, the use of this option can provide better distribution of incoming datagrams to multiple processes (or threads) as compared to the traditional technique of having multiple processes compete to receive datagrams on the same socket.


### 0x02 SO_REUSEPORT的问题：
1. 查找的时间复杂度伪O(N): 4.6已解决
    - [Linux 4.6内核对TCP REUSEPORT的优化](https://blog.csdn.net/dog250/article/details/51510823) (dog250)
2. 无法实现哈希一致性
    - [重新实现reuseport逻辑，实现一致性哈希](https://blog.csdn.net/dog250/article/details/89268404) (dog250)

3. 多队列阻塞问题
    - [从SO_REUSEPORT服务器的一个弊端看多队列服务模型](https://blog.csdn.net/dog250/article/details/107227145) (dog250)

即使`SO_REUSEPORT`不是完美的，但是大多数情况下，使用它会是个明智的选择。`nginx`在[Socket Sharding in NGINX Release 1.9.1](https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/) 中的测试表明，使用`SO_REUSEPORT`之后，服务在延时、延时方差、以及性能这三个方面都有了非常明显的改进。

### 0x03 SO_REUSEPORT与SO_REUSEADDR的区别
- [How do SO_REUSEADDR and SO_REUSEPORT differ?](https://stackoverflow.com/questions/14388706/ how-do-so-reuseaddr-and-so-reuseport-differ) (stackoverflow)


### 0x04 其他用途——seemless reloads
- [GLB part 2: HAProxy zero-downtime, zero-delay reloads with multibinder](https://github.blog/2016-12-01-glb-part-2-haproxy-zero-downtime-zero-delay-reloads-with-multibinder/)


