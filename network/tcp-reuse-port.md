## SO_REUSEPORT备忘

### 0x00 epoll的惊群问题
为了理解`SO_REUSEPORT`，有必要先了解一下`epoll`的`惊群问题`
- [Epoll is fundamentally broken 1/2](https://idea.popcount.org/2017-02-20-epoll-is-fundamentally-broken-12/)
再看一下cloudfare这篇[Why does one NGINX worker take all the load?](https://blog.cloudflare.com/the-sad-state-of-linux-socket-balancing/)

### 0x01 SO—REUSEPORT是什么


- 在学习`SO_REUSEPORT`的过程中，我们应该先读一遍LWN上[The SO_REUSEPORT socket option](https://lwn.net/Articles/542629/)，这篇文章也可以看作是对这一特性的官方介绍。

- 掘金小册`《深入理解 TCP 协议：从原理到实战》`中的[一台主机上两个进程可以同时监听同一个端口吗](https://juejin.im/book/6844733788681928712/section/6844733788832923661)对这个问题有很全面的讲解，也值得一读。


### 0x02 SO_REUSEPORT的问题：
1. 查找的时间复杂度伪O(N)
    - [Linux 4.6内核对TCP REUSEPORT的优化](https://blog.csdn.net/dog250/article/details/51510823)
2. 无法实现哈希一致性
    - [重新实现reuseport逻辑，实现一致性哈希](https://blog.csdn.net/dog250/article/details/89268404) (dog250大佬的文章)


### 0x03 SO_REUSEPORT与SO_REUSEADDR的区别
- 请阅读stackoverflow上的这篇[How do SO_REUSEADDR and SO_REUSEPORT differ?](https://stackoverflow.com/questions/14388706/how-do-so-reuseaddr-and-so-reuseport-differ)





