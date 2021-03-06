## 什么是vxlan



> VXLAN 全称是 Virtual eXtensible Local Area Network，虚拟可扩展的局域网。它是一种 overlay 技术，通过三层的网络来搭建虚拟的二层网络

这句话怎么理解？

首先，分析句子主干，可以知道vxlan的目的是构建一个二层网络？

那么二层网络有什么特点呢？一个广播域📢！

也就是说，通信双方认为自己在一个局域网内，通信行为**看起来**与真实局域网内的机器间通信别无二致。

实际上真的需要各个主机在同一个物理局域网内么？not really!

再看这句话，目的是构建二层网络，手段是通过三层网络。
三层网络啥概念？就是联网的概念。比如看这条微博的你，和我就是三层可通的，而我们不在同一个网络(指局域网)。

那这种网络到底是怎么构建的？

通过封包与解包为主要手段的隧道技术。

![vxlan图示](./vnet-vxlan.png)

简单来说，就是把我们要通信的原始报文(我们以为我们在一个局域网内通讯，比如我172.16.1.2/24想要找172.16.1.3/24)，外层加上一层vxlan header, 整体作为payload，通过UDP协议通信。所以“原始报文+vxlan header”是UDP通信的payload, 这个UDP通信自然是利用三层网络了，和我们平时网上冲浪🏄的过程一样。


## 推荐阅读
- [vxlan协议原理](https://cizixs.com/2017/09/25/vxlan-protocol-introduction)
- [linux 上实现 vxlan 网络](https://cizixs.com/2017/09/28/linux-vxlan/)