## 问题
    想在myql中添加一个账号，并且授予其某些数据库的所有权限。

## 解决办法

1. 先用`root`账户登录，查看已有的用户:
    ```mysql
    mysql> select user from mysql.user;

    +------------------+
    | user             |
    +------------------+
    | mysql.infoschema |
    | mysql.session    |
    | mysql.sys        |
    | root             |
    +------------------+
    4 rows in set (0.00 sec)
    ```

2. 创建新用户

    ```mysql
    mysql> create user lfn@localhost identified by 'Secure1pass!';
    ```
    再次查看数据库中用户:

    ```mysql
    mysql> select user from mysql.user;
    +------------------+
    | user             |
    +------------------+
    | lfn              |
    | mysql.infoschema |
    | mysql.session    |
    | mysql.sys        |
    | root             |
    +------------------+
    5 rows in set (0.00 sec)
    ```
    可以看到，用户`lfn`已经如期添加。

3. 新建数据库

    ```mysql
    mysql> create database "example";
    ```

4. 授权
    ```mysql
    mysql> grant all privileges on example.* to lfn@localhost;
    ```

5. 确认

    以用户名`lfn@localhost`登录，确认能看到`example`数据库。
    ```bash
    ~ mysql -ulfn -p
    ```

    ```mysql
    mysql> show databases;
    +--------------------+
    | Database           |
    +--------------------+
    | information_schema |
    | example            |
    +--------------------+
    2 rows in set (0.00 sec)

    ```

## 推荐阅读
- [How To Create User Accounts Using MySQL CREATE USER Statement](https://www.mysqltutorial.org/mysql-create-user.aspx)