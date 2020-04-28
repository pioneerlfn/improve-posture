## TCP的计时器(timer)

这个话题比较有意思，值得详细总结.看过之后能对于各个计时器超时时间的大小有个直观认识.

TCP 为每条连接建立了 7 个定时器：
1. 连接建立定时器

    建立连接时，最多将重传 6 次（间隔 1s、2s、4s、8s、16s、32s），6 次重试以后放弃重试，connect 调用返回 -1，调用超时.所以这个超时时间默认是1min左右;

2. 重传定时器

    重传时间间隔是指数级退避，直到达到 120s 为止，重传次数是15次（这个值由操作系统的 /proc/sys/net/ipv4/tcp_retries2 决定)，总时间将近 15 分钟。

3. 延迟 ACK 定时器

    在 TCP 收到数据包以后在没有数据包要回复时，不马上回复 ACK。这时开启一个定时器，等待一段时间看是否有数据需要回复。如果期间有数据要回复，则在回复的数据中捎带 ACK，如果时间到了也没有数据要发送，则也发送 ACK。在 Centos7 上这个值为 40ms(注意是毫秒！)

4. PERSIST 定时器

    Persist 定时器是专门为零窗口探测而准备的。我们都知道 TCP 利用滑动窗口来实现流量控制，当接收端 B 接收窗口为 0 时，发送端 A 此时不能再发送数据，发送端此时开启 Persist 定时器，超时后发送一个特殊的报文给接收端看对方窗口是否已经恢复，这个特殊的报文只有一个字节。重试的策略跟前面介绍的超时重传的机制一样，时间间隔遵循指数级退避，最大时间间隔为 120s，重试了 16，总共花费了 16 分钟

5. KEEPALIVE 定时器

    tcp连接建立之后，如果不主动关闭，将一直保持连接ESTABLESHED状态。为了检测上层已经失效的连接，默认大概2h会发探活包，也即tcp的keep-alive.注意这个时间大概是2h.但是我们知道网路连接是一种relay的形式，中间会有层次路由或者代理，一般的连接不用2h,中间的路由或者代理会关闭不活跃的连接。有时候为了维持不被中间层擅自关闭，需要应用层主动探活。

6. FIN_WAIT_2 定时器

    四次挥手过程中，主动关闭的一方收到 ACK 以后从 `FIN_WAIT_1` 进入 `FIN_WAIT_2` 状态等待对端的 FIN 包的到来，`FIN_WAIT_2` 定时器的作用是防止对方一直不发送 FIN 包，防止自己一直傻等。这个值由`/proc/sys/net/ipv4/tcp_fin_timeout` 决定，在 Centos7 机器上，这个值为 `60s`，也即1min左右。与握手阶段那个超时时间相仿;
    
    关于`FIN_WAIT_2`和`CLOSE_WAIT`出现的场景，可以看cloudflare的这篇博客:
    
    > [This is strictly a violation of the TCP specification](https://blog.cloudflare.com/this-is-strictly-a-violation-of-the-tcp-specification/)

7. TIME_WAIT 定时器

    大名鼎鼎的TIME_WAIT,面试最喜欢问TIME_WAIT 定时器也称为 2MSL 定时器，可能是这七个里面名气最大的，主动关闭连接的一方在 TIME_WAIT 持续 2 个 MSL 的时间，超时后端口号可被安全的重用.在Linux中,MSL被定义成30s, 所以2MSL大概是1min.                                 