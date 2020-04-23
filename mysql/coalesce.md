## 使用函数COALESCE 替换记录中的NULL值

在我们使用Go操作数据库的过程中，如果record中存在`NULL`值，代码会显得ugly.

一种解决办法是使用`database/mysql`中定义的NULL系列类型，比如`sql.NULLString`,我之前在代码中就是这么干的。

还有一种办法是使用MySQL的`COALESCE`函数，语法是这样的：

```
COALESCE(value1,value2,...);
```
`COALESCE`接受一堆参数，返回🔙第一个不为NULL的值。
因此我们的sql可以这么写:

```
SELECT 
COALESCE(name,'') 
FROM user 
WHERE id = 1
```
这样，如果name是`NULL`的话，就会返回空字符串`''`，这样我们在后续处理比如调用`sql.Rows.Scan()`的时候就不会报错了。

## 推荐阅读
- [Introduction to MySQL COALESCE function](https://www.mysqltutorial.org/mysql-coalesce/)
- [Go database/sql tutorial: Working with NULLs](http://go-database-sql.org/nulls.html)