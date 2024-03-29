## 学习材料

- [Writing An Os In Rust](https://os.phil-opp.com/)
- [rCore-Tutorial-Book 第三版](https://rcore-os.github.io/rCore-Tutorial-Book-v3/index.html)
- [jyywiki.cn](http://jyywiki.cn/)
- [现代操作系统：原理与实现(银杏书)](https://ipads.se.sjtu.edu.cn/mospi/)
- [南京大学 计算机科学与技术系 计算机系统基础 课程实验 2021](https://nju-projectn.github.io/ics-pa-gitbook/ics2021/)
- [The Art Of Command Line](https://github.com/jlevy/the-art-of-command-line)
- [linux inside(中文版)](https://github.com/MintCN/linux-insides-zh)
- [深入理解计算机系统合集/csapp-概念讲解+OS代码实现](https://www.bilibili.com/video/BV17K4y1N7Q2?spm_id_from=333.999.0.0)
- [Linux 内核 0.12 完全注释-赵炯博士](http://www.oldlinux.org/download/CLK-5.0-WithCover.pdf)
- [你管这破玩意叫操作系统源码](https://github.com/sunym1993/flash-linux0.11-talk)
- [The Missing Semester of Your CS Education](https://missing-semester-cn.github.io/)
- [System V Application Binary Interface](http://jyywiki.cn/pages/OS/manuals/sysv-abi.pdf)
- ostep
- csapp
- ulk
- apue
- tlpi

## manual need to read
- [ ] sh (man sh, `dash` not `bash`)
- [ ] /proc文件系统
- [ ] mmap (ld大量使用/文件映射,缺页中断再读盘)
- [ ] fork (调用一次，返回2次，父子进程返回顺序不确定)
- [ ] exec (调用一次，用不返回 ==means=> reset状态机)
- [ ] syscall (相比int0x80更加优化)
- [ ] vsdo (不陷入kernel的系统调用，有趣的把戏, 比如gettimeofday())
- [ ] posix_spawn
- [ ] [A fork() in the road](https://www.microsoft.com/en-us/research/uploads/prod/2019/04/fork-hotos19.pdf)
