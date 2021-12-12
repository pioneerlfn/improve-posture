## 前言
wget是经典的下载工具，我们平时可能只用很少的几个参数。
但是如果你用wget来下载整个站点，那你可能需要掌握更多的参数。<br><br>


## 参数详解
> todo: TBD
- `-N`
    ```shell
    -N 代表turn on timestamp
    等价于 --timestamping
    ```
- `-c`
- `-nv`
- `--retry-connrefused`
- `--waitretry=1`
- `-e robots=off`
- `--user=xxx`
- `--password=yyy`
- `-timeout=10`
- `-t 3`
- `--limit-rate`
- `-P`
- `--cut-dirs`
- `--preserve-permissions`
- `-r`
- `-l inf`
- `-nH`
- `--reject=index.html*`

<br>

## 参考
- [GNU Wget 1.21.1-dirty Manual](https://www.gnu.org/software/wget/manual/wget.html#Time_002dStamping)
- [wget vs curl](https://daniel.haxx.se/docs/curl-vs-wget.html)