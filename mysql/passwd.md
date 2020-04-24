## 重置mysql的root密码

### 环境
> OS: ubuntu 20.04

> mysql/mysqld: 8.0.19

### 问题
在ubuntu系统中，通常都是直接执行下面👇这行命令安装`mysql`服务端和客户端:
```bash
sudo apt install mysql-client mysql-server -y
```

安装之后尝试登陆，发现我们不知道`root`的密码，安装过程中也没有提示我们设置，不知道是不是因为加了`-y`参数的原因。

不过，也不用慌。

### 解决办法

除了root用户，我们还可以用其他用户登陆，而且密码已知：

打开`/etc/mysql/debian.cnf`，我们可以看到这样的内容：

```
[client]
host  = localhost
user  = debian-sys-maint
password = xxxxxxx
socket  = /var/run/mysqld/mysqld.sock
```

试一下用上面的`debian-sys-maint`，发现可以成功登录。然后用执行:
```mysql
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'yourpassword'
```
退出，重新以`root`账户登录即可。

> 新用户

> 创建新用户，可以执行下面这句:
>```mysql
>create user lfn@localhost identified by >'Secure1pass!';
>```