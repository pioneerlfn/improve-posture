## 问题

给一个英文文本文件，输出前十个出现次数最多的单词


## 可行脚本

```shell

cat <filename> | tr -s '[:space:]' '\n' |tr '[:upper:]' '[:lower:]'|sort|uniq -c|sort -nr|head -10 

```

## 命令解释

### cat

> concatenate files and print on the standard output(连接文件，并且打到标准输出)

举例：
```bash

➜  shell echo "hello world" > 1; echo "hello lfn" >  2

➜  shell cat 1 2
hello world
hello lfn

➜  shell cat 2 1
hello lfn
hello world
```

### tr

> translate or delete characters(删除或转换字符)

    -s, --squeeze-repeats
    replace  each  input  sequence  of  a  repeated  character  that  is  listed in SET1 with a single occurrence of that character

在`tr -s '[:space:]' '\n'`, 多个空格会被替换成一个回车符。

### sort

> sort lines of text files(将文件行排序)
    
    -r
    倒序排序

### uniq -c

> uniq: report or omit repeated lines

> -c, --count
       prefix lines by the number of occurrences

`uniq -c`会将统计重复行的数目，并且写到行前.

### head -10

输出前10行
