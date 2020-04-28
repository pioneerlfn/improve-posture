## FIN_WATI_2与CLOSE_WAIT

这篇文章可以回答这样一个问题:
> 如果连接的一方主动关闭连接之后(记为client)，被动关闭方未能调用close关闭连接(记为server)，双方分别会处在什么状态？

回顾一下事件序:

1. client 调用`close`之后,发送`FIN`给对端，自己状态转为`FIN_WAIT_1`;

2. server 收到`FIN`之后，发送`ACK`给 client, 自己状态转为`CLOSE_WAIT`
    > CLOSE_WAIT实际就是wait close的意思。

3. client收到`ACK`, 自己转为`FIN_WAIT_2`, 同时启用一个计时器，默认超时时间为60s.


看一下开头我们提的问题，如果此时server端没能调用`close`关闭socket，会出现的情况是:

1. 对于client, 由于启用了计时器，所以不会一直傻等，超时时间之后(60s),内核会销毁该 socket, 回收端口;
2. 对于server, 由于一直不调用`close`, 所以对应的socket一直处于`CLOSE_WAIT`, 会造成socket的泄露

这里除了服务端资源泄露之外，其实还存在其他的隐患。cloudflare的这篇 [This is strictly a violation of the TCP specification](https://blog.cloudflare.com/this-is-strictly-a-violation-of-the-tcp-specification/) 提供了一个例子：

假设server 在5000端口listen, client与server建立连接，client kernerl为连接选的端口是X,那么在双方机器上的连接就是下面👇这样的:

- client: (X --> 5000)
- server: (5000 --> X)

经历前面分析的泄露过程之后, 经过一段时间，client回收了端口X，而server上的连接仍然未释放:
- server: (5000 --> X), `CLOSE_EAIT`

此时，client多次连接server, 如果client 的kernel又选了X这个端口来连接server, 那么:
- client: (X --> 5000), `SYN_SENT`
- server: (5000 --> X), `CLOSE_WAIT`

由于在server端，(5000 --> X)处于`CLOSE_WAIT`状态，所以server的kernel就困惑了，是不会响应client的这次握手🤝请求的，client最终超时，连接失败。

## 推荐阅读
本来以为`FIN_WAIT_2`比较简单，看了下面两篇文章，发现还是很琐碎复杂的:
- [TCP 之 FIN_WAIT_2状态处理流程](http://www.linuxtcpipstack.com/537.html)
- [TCP套接口的FIN_WAIT_2状态超时](https://blog.csdn.net/sinat_20184565/article/details/88562876)


