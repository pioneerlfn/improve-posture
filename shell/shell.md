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

## sort
### 按某列排序: `sort -k` 以及统计相同记录出现的次数：`uniq -c`

```shell
k get po -owide | grep spread-ks-225b6-7263-x8

spread-ks-225b6-7263-x8-deployment-6747dfc94-4xhbh                3/3     Running             0          12m     10.133.37.48    10.86.114.11   <none>           <none>
spread-ks-225b6-7263-x8-deployment-6747dfc94-5dzwc                3/3     Running             0          12m     10.133.37.28    10.86.142.39   <none>           <none>
spread-ks-225b6-7263-x8-deployment-6747dfc94-7bh54                3/3     Running             0          12m     10.133.36.202   10.86.100.45   <none>           <none>
spread-ks-225b6-7263-x8-deployment-6747dfc94-8wpgr                3/3     Running             0          12m     10.133.37.40    10.86.98.32    <none>           <none>
spread-ks-225b6-7263-x8-deployment-6747dfc94-mqdqq                3/3     Running             0          12m     10.133.37.51    10.86.114.11   <none>           <none>
spread-ks-225b6-7263-x8-deployment-6747dfc94-nstdv                3/3     Running             0          12m     10.133.37.60    10.86.98.32    <none>           <none>
spread-ks-225b6-7263-x8-deployment-6747dfc94-nsxc6                3/3     Running             0          12m     10.133.37.39    10.86.100.45   <none>           <none>
spread-ks-225b6-7263-x8-deployment-6747dfc94-rcjtk                3/3     Running             0          12m     10.133.37.46    10.86.114.11   <none>           <none>
spread-ks-225b6-7263-x8-deployment-6747dfc94-xbvlz                3/3     Running             0          12m     10.133.37.38    10.86.142.39   <none>           <none>
spread-ks-225b6-7263-x8-deployment-6747dfc94-xpvt8                3/3     Running             0          12m     10.133.37.68    10.32.19.35    <none>
```
其中，第七列是`宿主机ip`, 为了找出pod在宿主机中的分布情况，可以用下列命令:

```shell
k get po -owide | grep spread-ks-225b6-7263-x8 | sort -k 7  | awk '{print $7}' | uniq  -c | sort

      1 10.32.19.35
      2 10.86.100.45
      2 10.86.142.39
      2 10.86.98.32
      3 10.86.114.11
```
可以看出，这个deployment的pod在宿主机间的分布式不均匀的。
> Note: 上面👆例子中可以看出`uniq`的用法，`-c`用来统计重复记录出现的次数，也即 Display number of occurrences of each line along with that line:

## sed
### 删除空格 `sed s/[[:space:]]//g`

## comm
### 找出两个个文件中相同的记录 `comm -12 file1 file2`
> 注意：使用`comm`前需要先用`sort`排序


## 生成token
```shell
head -c 16 /dev/urandom | od -An -t x | tr -d ' '

```
> /dev/urandom
> https://unix.stackexchange.com/questions/324209/when-to-use-dev-random-vs-dev-urandom

## 对比目录差量
```
shell

rsync -ainv --delay-updates --ignore-times --hard-links
```
> 解释见https://explainshell.com/explain?cmd=rsync+-ainv+--delay-updates+--ignore-times+--hard-links+

## tar
```shell
	-a # 根据提供的包名自动压缩, 比如xxx.tgz会自动用gzip压缩

```

## tty
- [TTY 到底是个什么玩意](https://mp.weixin.qq.com/s/MNJ3DrXBkzgwxE8yTBCRBQ?v_p=90&WBAPIAnalysisOriUICodes=10000011_10000011&launchid=10000365--x&wm=3333_2001&aid=01A0pWWb1rXSEKlcJ1hBqP4K0wekuzaAJ0XJ-7LMHEmjjb5fA.&from=10C3593010)
- [reverse shell cheetsheet](https://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet)













