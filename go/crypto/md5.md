## 计算字符串的MD5值

- MD5消息摘要算法（英语：MD5 Message-Digest Algorithm），一种被广泛使用的密码散列函数
- 可以产生出一个`128位（16字节`）的散列值（hash value）
- 用于确保`信息传输完整一致`
- 一般128位的MD5散列被表示为32位十六进制数字

```Go
package main

import (
    "crypto/md5"
    "fmt"
)

func main() {
    data := []byte("hello")
    // 用md5.sum直接计算
    // 然后以16进制形式输出
    fmt.Printf("0x%x\n", md5.Sum(data)) // 0x5d41402abc4b2a76b9719d911017c592
    
    // ""空字符串的md5
    fmt.Printf("0x%x\n", md5.Sum([]byte(""))) // 0xd41d8cd98f00b204e9800998ecf8427e
}

```

## 推荐阅读
- [How to get a MD5 hash from a string in Golang?](https://stackoverflow.com/questions/2377881/how-to-get-a-md5-hash-from-a-string-in-golang)
- [wikipedia: MD5](https://zh.wikipedia.org/wiki/MD5)