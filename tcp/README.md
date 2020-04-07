梨叔上个月发了条微博，问了这样[一个问题](https://m.weibo.cn/status/4484226152273449?sudaref=login.sina.com.cn),我试着回答一下。


有两种方式可以关闭 TCP 连接

- FIN：优雅关闭，发送 FIN 包表示自己这端所有的数据都已经发送出去了，后面不会再发送数据
- RST：强制连接重置关闭，无法做出什么保证


## SO_LINGER
Linux 的套接字选项SO_LINGER 用来改变socket 执行 close() 函数时的默认行为。


SO_LINGER 启用时，操作系统开启一个定时器，在定时器期间内发送数据，定时时间到直接 RST 连接。

SO_LINGER 参数是一个 linger 结构体，代码如下

```
struct linger {
    int l_onoff;    /* linger active */
    int l_linger;   /* how many seconds to linger for */
};
```

第一个字段 l_onoff 用来表示是否启用 linger 特性，非 0 为启用，0 为禁用 ，linux 内核默认为禁用。这种情况下 close 函数立即返回，操作系统负责把缓冲队列中的数据全部发送至对端

第二个参数 l_linger 在 l_onoff 为非 0 （即启用特性）时才会生效。

如果 l_linger 的值为 0，那么调用 close，close 函数会立即返回，同时丢弃缓冲区内所有数据并立即发送 RST 包重置连接
如果 l_linger 的值为非 0，那么此时 close 函数在阻塞直到 l_linger 时间超时或者数据发送完毕，发送队列在超时时间段内继续尝试发送，如果发送完成则皆大欢喜，超时则直接丢弃缓冲区内容 并 RST 掉连接。

**我们进假设一般没有人会去调这个flag.**
因此接下来的讨论将不考虑这种情况。


## OS panic（or poweroff）/ 网络故障

这时候OS突然崩溃或者断电，来不及向网络连接发送任何信息，所以对端此时还无法感知，还在buffer中未发送的数据也一并丢失。如果应用程序不发送数据，可能永远无法得知该连接已经失效。


## 正常结束 or 异常结束？ No Matter！

如果OS不崩溃，进程无论是正常还是异常退出，内核处理最后会走到`do_exit`函数，在`do_exit`会调用`exit_files`释放文件对象。

> All of the **file descriptors**, directory streams, conversion descriptors, and message catalog descriptors open in the calling process shall be **closed**.

所以在OS不崩溃的情况下，不论进程是正常还是异常退出，OS都会关闭打开的socket, 发送`FIN`给对端socket.


## 对端如何感知？

1. 读
    - **服务器OS未奔溃，可有立即知道**

        如果客户端进程正阻塞在read调用上（前面已经提到，此时receive buffer一定为空，因为read在receive buffer有内容时就会返回），则read调用立即返回EOF，进程a被唤醒，得知连接已关闭。

    - **服务器OS崩溃，大概只能傻等**

        客户端无感知，如果无特殊处理逻辑，那只能等TCP的`keep_alive`超时机制发挥作用。TCP 协议的设计者考虑到了这种检测长时间死连接的需求，于是乎设计了 keepalive 机制。一般来说keepalive会等2小时(7200s)发送探测包，探测 9 次，每次探测间隔 75s，这些值都有对应的参数可以配置。

2. 写

    等客户端有数据有数据要发送给服务端时，服务端这边并没有这条连接的信息，发送 RST 给客户端，告知客户端自己无法处理。

    这需要客户段两次写才能发现。第一次 write 会触发服务端发送 RST 包，这时客户端内核会知道连接已经无效，用户程序还不知道，等用户程序第二次 write 时会抛出Broken pipe异常，从而得知连接已断开。这是因为：

    > “当一个进程向某个已收到 RST 的套接字执行写操作时，内核向该进程发送一个 SIGPIPE 信号。该信号的默认行为是终止进程，因此进程一般会捕获这个信号进行处理。不论该进程是捕获了该信号并从其信号处理函数返回，还是简单地忽略该信号，写操作都将返回 EPIPE 错误（也就Broken pipe 错误）”





