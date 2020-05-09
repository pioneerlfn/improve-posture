## Pod状态速查

- Pending

    如果一个pod卡在Pending状态，则表示这个pod没有被调度到一个节点上
- Waiting

    如果一个pod卡在Waiting状态，则表示这个pod已经调试到节点上，但是没有运行起来。
    
     再次敲一下kubectl describe ...这个命令来查看相关信息。 最常见的原因是拉取镜像失败。

- 