# improve-posture

一个后端程序员的自我修养。

## mysql

- [彻底解决MySQL中的乱码问题](https://mp.weixin.qq.com/s/58Y11c8cLN1uDfHn_6lyAg)
- [利用函数COALESCE 处理记录中的NULL](./mysql/coalesce.md)

## redis


- [理解Redis的内存回收机制](https://juejin.im/post/5d107ad851882576df7fba9e?utm_source=weibo&utm_campaign=user)


## 网络篇

- [TCP连接的若干细节](./tcp/README.md)
- [怎么理解vxlan](./network/vxlan.md)
- [ARP:一个很简单的协议](./network/arp.md)


## web

- [web application vs web service](./web/app-service.md)
- [HTTP Method](./web/http.md)
- [非对称加密 && 数字签名:rsa实战](./go/crypto/rsa.go)
- [jwt实战: token生成与解析](./jwt/README.md)
- OAuth
    - [How to do Google sign-in with Go](https://skarlso.github.io/2016/06/12/google-signin-with-go/)
    - [How to do Google Sign-In with Go - Part 2](https://skarlso.github.io/2016/11/02/google-signin-with-go-part2/)


## 运维篇

- [mysql添加用户并收入数据库权限](./mysql/user/README.md)
- [mysql登录问题](./mysql/login/README.md)
- [ubuntu 20.04安装Mysql后root密码未知](./mysql/passwd.md)
- [Sequel Pro无法刷出数据的问题](./tools/sequelpro.md)
- [在终端pprint json](./tools/json/print.md)
- [ssh反向端口转发采坑](./tools/ssh.md)
- [听说你也要变基](./tools/rebase.md)

## 语言篇(Go)

- [Go中使用基于etcd的分布式锁](./go/distributed/locks.go)
- [defer相关问题](./go/defer/README.md)
- [为何这里不需要atomic?](./go/atomic/READEME.md)
- [Go学习笔记——测试技巧备忘](./go/testing/README.md)
- [如何多快好省拼接字符串?](./go/strings/README.md)
- [分隔字符串注意事项](./go/strings/split.md)
- [时区问题](./go/time/README.md)
- [写ResponseWriter的一点注意事项](./go/http/README.md)
- [使用EqualFold比价字符串大小](./go/strings/equalfold.md)
- [http.Request备忘](./go/http/request.md)
- [Unmarshal JSON](./go/json/unmarshal.md)
- [计算字符串的MD5散列值](./go/crypto/md5.md)
- [interface与反射使用示例:一个较为通用的数据库查询函数](./go/reflect/mysql.go)
- [并发执行任务:一段常用的代码](./go/concurrency/parralize.go)

## 语言篇(C)

- [setjmp-longjmp](./c/setjmp-longjmp.md)
- [注释pause.c(kubernetes)](./c/pause.c)
- [mmap初体验](./c/mmp.c)


## 语言篇(shell)

- [给一个英文文本文件，输出前十个出现次数最多的单词](./shell/top10.md)

## OS篇

- [初探cgroup](./os/cgroup.md)
- [cgroup之cpu子系统](./os/cgroup-cpu.md)