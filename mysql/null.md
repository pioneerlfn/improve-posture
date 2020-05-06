## 关于mysql中的NOT NULL，你需要知道的


`NOT NULL`是对列的一种约束(constraint).
1. `NOT NULL`一般在创建表的时候声明:
    ```sql
    column_name data_type NOT NULL;
    ```
2. 如果表中某列声明了`NOT NULL`, 则在`INSERT`或`UPDATE`的时候，该列的值不能为`NULL`, 否则会插入失败报错。

3. 声明为主键(`PRIMARY KEY`)的列会隐式包含`NOT NULL`约束。

4. 关于如何给已经存在的非 `NOT NULL` 列添加`NOT NULL`约束，或者去掉列上的`NOT NULL`约束，请参考 `mysqltutorial` 这篇 [MySQL NOT NULL Constraint.](https://www.mysqltutorial.org/mysql-not-null-constraint/)

5. 如果某列的值为 `NULL`, 则在  `innodb` 记录中，会用`NULL值列表`中对应的标志位表示，在记录的实际数据中并不存储。具体请参考小青蛙老师这篇 [InnoDB记录存储结构](https://juejin.im/book/5bffcbc9f265da614b11b731/section/5bffda656fb9a049b13deba8)

6. 为了测试某列是否为 `NULL`, 我们会用 `IS NULL`和`IS NOT NULL` 运算符:
    ```sql
    mysql> SELECT 1 IS NULL, 1 IS NOT NULL;
    +-----------+---------------+
    | 1 IS NULL | 1 IS NOT NULL |
    +-----------+---------------+
    |         0 |             1 |
    +-----------+---------------+
    ```

7. 任何数学运算符与 `NULL` 比较，结果都是 `NULL`, 得到的结果是无意义的. 请阅读官方文档 [4.4.6 Working with NULL Values](https://dev.mysql.com/doc/mysql-tutorial-excerpt/5.7/en/working-with-null.html)
    ```sql
    mysql> SELECT 1 = NULL, 1 <> NULL, 1 < NULL, 1 > NULL;
    +----------+-----------+----------+----------+
    | 1 = NULL | 1 <> NULL | 1 < NULL | 1 > NULL |
    +----------+-----------+----------+----------+
    |     NULL |      NULL |     NULL |     NULL |
    +----------+-----------+----------+----------+
    ```

8. 在mysql中，`0 或者 NULL` 代表 false, 其他所有值都代表 true. 布尔操作的默认真值是 `1`

9. 两个NULL值在`GROUP BY`操作中被认为是相等的.

10. 列中的值排序的时候, `NULL`被认为是最小的：
    > We define the SQL null to be the smallest possible value of a field.

10. `WHERE` 语句中出现 `IS NULL`, `IS NOT NULL`, `!=` 等比较符的时候，仍然可以用索引。感兴趣可以查看这篇 [MySQL中IS NULL、IS NOT NULL、!=不能用索引？胡扯！](https://juejin.im/post/5d5defc2518825591523a1db)

11. 如果联合索引中出现`NULL`值，会存在一些比较微妙的现象，感兴趣可以看下这篇 [谜之 NULL](https://blog.wolfogre.com/posts/sql-tips/#%E8%B0%9C%E4%B9%8B-null)

12. 除非有充足理由，否则请将每一列都声明为 `NOT NULL`.



## 推荐阅读
- [MySQL NOT NULL Constraint](https://www.mysqltutorial.org/mysql-not-null-constraint/)
- [4.4.6 Working with NULL Values](https://dev.mysql.com/doc/mysql-tutorial-excerpt/5.7/en/working-with-null.html)
- [InnoDB记录存储结构](https://juejin.im/book/5bffcbc9f265da614b11b731/section/5bffda656fb9a049b13deba8)
- [谜之 NULL](https://blog.wolfogre.com/posts/sql-tips/#%E8%B0%9C%E4%B9%8B-null)
- [MySQL中IS NULL、IS NOT NULL、!=不能用索引？胡扯！](https://juejin.im/post/5d5defc2518825591523a1db)

