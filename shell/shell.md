## 给一个英文文本文件，输出前十个出现次数最多的单词
可行脚本

```shell
cat <filename> | tr -s '[:space:]' '\n' |tr '[:upper:]' '[:lower:]'|sort|uniq -c|sort -nr|head -10 
```

## 项目中查找包含特定内容的行
```shell
grep -n `find 目录1 目录2 目录3 -name "*.go"` -e "createTaskMonitorData"
```
这行命令的意思，是在目录1，目录2，目录3下的`go` 文件中查找包含`createTaskMonitorData`的行.

这行命令特别常用。

## git
pretty print git log:
```shell
git log --all --decorate --oneline --graph
```