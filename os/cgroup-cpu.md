## Cgroup之cpu子系统

cpu 子系统可以调度 cgroup 对 CPU 的获取量。可用以下两个调度程序来管理对 CPU 资源的获取：

- 完全公平调度程序（CFS） — 一个比例分配调度程序，可根据任务优先级 ∕ 权重或 cgroup 分得的份额，在任务群组（cgroups）间按比例分配 CPU 时间（CPU 带宽）
- 实时调度程序（RT） — 一个任务调度程序，可对实时任务使用 CPU 的时间进行限定。

### CFS
对cpu资源的限制有2中，一种是绝对值，一种是相对值。
1. 绝对值

    绝对值强制对CPU的使用不超过一定的上限。需要配置下面2个参数
    - cpu.cfs_period_us
        1. `cfs_period_us`表示对1个cpu的分割间隔，单位是us
        2. 上限为1s(1000000), 下限为1000us(1000)

    - cpu.cfs_quota_us

        `cfs_quata_us`根据与`cfs_period_us`的比例来表示占用多少的CPU资源。
    
        比如`cfs_period_us`为250ms(250000), `cfs_quota_us`为500ms(500000),那就代表2个CPU，如果为100ms, 那就是1/4个cpu资源。
2. 相对值
    
    相对值由参数`cpu_shares`确定。

    顾名思义，这个参数代表的是占比，虽然对每个cgroup设定的是一个绝对的数值，但是获得的CPU比例，要按照这个数值占所有cgroup这个数值的总和的比例来定。

    举例：
    
    - Cgroup-A: cpu_shares=100
    - Cgroup-B: cpu_shares=100
    
    那么Cgroup-A与Cgroup-B都可以占用50%的CPU资源。
    
    此时如果新加入一个Cgroup-C:
    - Cgroup-C: 200
    
    那么：
    - Cgropu-A: 25%
    - Cgropu-B: 25%
    - Cgropu-C: 50%

    需要注意，`cpu_shares`不对CPU的使用构成硬约束，比如Cgroup-C暂时没用到那么多CPU资源，那么A和B是可以超过25%的。

    



### RT
实时进程调度与CFS类似，不再赘述。