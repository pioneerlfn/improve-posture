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

### squash多次提交
同一个功能或者同一篇文章，我们可能多次提交，导致git log看起来凌乱琐碎。
在push到远端之前，可以通过`git rebase -i`来压缩多次提交。
1. 执行下面👇这行命令，就可以查看并合并最近N次的提交:

```bash
git rebase -i HEAD~N
```
2. 这时候，会自动进入 vi 编辑模式.注释显示，有下面这些动作:

- p, pick = use commit
- r, reword = use commit, but edit the commit message
- e, edit = use commit, but stop for amending
- s, squash = use commit, but meld into previous commit
- f, fixup = like "squash", but discard this commit's log message
- x, exec = run command (the rest of the line) using shell
- d, drop = remove commit

    要压缩commit的话，我们可以将对应提交前面的`pick`改成`s(squash)`,代表将本次提价与上次提交压缩到一起。

3. 修改完之后,`wq`退出。

    如无意外，提交记录📝应该看起来干净多了。
4. 如果保存的时候出现 `error: cannot 'squash' without a previous commit`导致退出了vi编辑窗口，执行:
    ```bash
    git rebase --edit-to
    ```
    重新进入编辑，改好之后保存退出。再执行
    ```bash
    git rebase --continue
    ```
    即可。

## 推荐阅读
- [彻底搞懂 Git-Rebase](http://jartto.wang/2018/12/11/git-rebase/)
- [3.6 Git 分支 - 变基](https://git-scm.com/book/zh/v2/Git-%E5%88%86%E6%94%AF-%E5%8F%98%E5%9F%BA)