## 前言
在目前的项目中，几乎没有见到过用连表查询的，所有的用例都是单表查询，不太清楚这是不是一种默认的公司规范。不过偶然也会有领导让你零时统计一些数据的需求，这个时候写程序统计可能略显笨重，而连表查询就有了用武之地。<br><br>

## 问题简述
有两张表A和B，统计既满足表A中某些字段的特征，又满足表B中某些字段特征的记录，对`id`去重。<br><br>

## sql语句
```sql
SELECT distinct(a.id) # 去重
FROM A_balana AS a
INNER JOIN B_yourlove AS b # 求交
ON a.id=b.id # 连表字段
WHERE a.group_info like '%docker%' # 表A的筛选条件
AND a.post_hook != 'xxxxx'   # 表A的筛选条件
AND b.last_time > '2021-10-30 18:00:00'   # 表B的筛选条件
GROUP BY a.id # 分组
ORDER BY a.id; # 排序
```
<br>

## 题外
我们很多人包括我自己背了各种各样的八股，对`ACID`倒背如流，但是却无法写出基本的`sql`语句。这是典型的眼高手低。让我们扎扎实实写代码，一点点变成一个合格的工程师吧。


