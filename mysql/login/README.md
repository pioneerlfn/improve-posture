## 问题
在mac上装了mysql之后，可以在命令行通过mysql客户端登录，但是无法通过`Sequel Pro`登录，报错信息如下：

`Authentication plugin 'caching_sha2_password' cannot be loaded: dlopen(/usr/local/lib/plugin/caching_sha2_password.so, 2): image not found`

## 环境
- OS: macOX 10.12
- mysql: 8.0.19


## 解决办法

在终端中通过mysql client以root身份登录，执行:

```mysql
ALTER USER 'lfn'@'localhost' IDENTIFIED WITH mysql_native_password BY 'yourpassword';
Query OK, 0 rows affected (0.01 sec)
```

再次登录`Sequel Pro`,无论是以`unix socket`还是`TCP`的方式，都可以成功登录。

## 参考文章
- [MySQL 中 localhost 127.0.0.1 区别](https://jin-yang.github.io/post/mysql-localhost-vs-127.0.0.1-introduce.html)
- [（解决方法）MySQL 8 + macOS 错误：Authentication plugin 'caching_sha2_password' cannot be loaded](https://1c7.me/mysql-8-connect-error-2018/)
