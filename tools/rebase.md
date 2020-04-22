# 听说你也要变基

感谢@游真uZ 和@老码农不上班 老师，重写一下这篇。

### 什么是变基

在 Git 中整合来自不同分支的修改主要有两种方法：`merge` 以及 `rebase`.

`rebase`就是变基。
```bash
git:($current-branch)  git rebase brach-x
```
上面这行命令，是把当前所在分支`current-branch`对应 `branch-x` 的新增，逐个加到以 `branch-x` 为基新开的分支上，添加完成后再把`current-brach`指向最新的commit piont，则`current-branch`现在同时包含了`branch-x`和`current-branch`的提交。

这种以`rebase`对象分支为基准的分支合并，导致`current-branch`的base由整合前与`branch-x`的最小公共节点，变成了`branch-x`的最新提交点，基(base)变了。


### 使用场景

通常有两种场景用 rebase 
1. 多人在`一个分支`协作，pull 代码时带上 `--rebase` 参数.

    这样拉去别人在协作分支时就不会产生一个 merge commit.  

2. feature 或 bugfix 分支同步其他分支最新提交。

    通常是 base 分支，比如`rebase` `dev`分支

    ```bash
    git:($bugfix) git rebase dev
    ```


## 推荐阅读
- [3.6 Git 分支 - 变基](https://git-scm.com/book/zh/v2/Git-%E5%88%86%E6%94%AF-%E5%8F%98%E5%9F%BA)