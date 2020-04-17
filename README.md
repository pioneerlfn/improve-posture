# improve-posture

一个后端程序员的自我修养。

## mysql

- [账户问题](./mysql/user/README.md)
- [登录问题](./mysql/login/README.md)

### 常用SQL

### 数据类型

- NULL

### 字符集(编码)

- utf8_general_ci
- utf8mb4
- [彻底解决MySQL中的乱码问题](https://mp.weixin.qq.com/s/58Y11c8cLN1uDfHn_6lyAg)



### 索引

- 覆盖索引
- 索引下推
- 联合索引
- 唯一索引

### 事务

## redis

### 原理

- 内存回收

    [理解Redis的内存回收机制
](https://juejin.im/post/5d107ad851882576df7fba9e?utm_source=weibo&utm_campaign=user)

### 实践

## kafka

## nginx



## TCP

- [TCP连接的若干细节](./tcp/README.md)


## web

- [web application vs web service](./web/app-service.md)

### HTTP

- [HTTP Method](./web/http.md)

- Authorization

- 非对称加密

- 数字签名

- jwt

    [jwt实战: token生成与解析](./jwt/README.md)
- OAuth
    - [How to do Google sign-in with Go](https://skarlso.github.io/2016/06/12/google-signin-with-go/)
    - [How to do Google Sign-In with Go - Part 2](https://skarlso.github.io/2016/11/02/google-signin-with-go-part2/)
- SSO(单点登录)
- cookie vs localStorage
- base64
- 跨域(CORS)
- 第三方登录

## 工具篇

### git

合并commits

### Sequel Pro

- [无法刷出数据的问题](./tools/sequelpro.md)

### json

- [在终端pprint json](./tools/json/print.md)


## 语言篇(Go)

### defer
- [defer相关问题](./go/defer/README.md)

### atomic

- [为何这里不需要atomic?](./go/atomic/READEME.md)

### go test

- [Go学习笔记——测试技巧备忘](./go/testing/README.md)

### 字符串

- [如何多快好省拼接字符串?](./go/strings/README.md)
- [分隔字符串注意事项](./go/strings/split.md)
### 时间问题

- [时区问题](./go/time/README.md)

### HTTP

- [写ResponseWriter的一点注意事项](./go/http/README.md)


## 语言篇(C)

- [setjmp-longjmp](./c/setjmp-longjmp.md)
- [注释pause.c(kubernetes)](./c/pause.c)


## 语言篇(shell)

- [给一个英文文本文件，输出前十个出现次数最多的单词](./shell/top10.md)