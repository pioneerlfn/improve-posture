## Cgroup子系统之cpuset

## 背景知识
请先阅读 [cgroup 之 cpuset 简介](https://jin-yang.github.io/post/linux-cgroup-cpuset-subsys-introduce.html)， 确保对cpuset有个基本的认知。

也就是说，`cpuset`赋予了我们绑核的能力。

在使用容器的时候，你可以通过设置 cpuset 把容器绑定到某个 CPU 的核上，而不是像 cpushare 那样共享 CPU 的计算能力。

这种情况下，由于操作系统在 CPU 之间进行上下文切换的次数大大减少，容器里应用的性能会得到大幅提升。

事实上，cpuset 方式，是生产环境里部署在线应用类型的 Pod 时，非常常用的一种方式。

## k8s中如何设置cpuset
其实非常简单。
- 首先，你的 Pod 必须是 Guaranteed 的 QoS 类型；
- 然后，你只需要将 Pod 的 CPU 资源的 requests 和 limits 设置为 `同一个相等的整数值` 即可.

比如,下面👇这个例子:

```yaml
spec:
  containers:
  - name: nginx
    image: nginx
    resources:
      limits:
        memory: "200Mi"
        cpu: "2"
      requests:
        memory: "200Mi"
        cpu: "2"
```

这时候，该 Pod 就会被绑定在 2 个独占的 CPU 核上。

当然，具体是哪两个 CPU 核，是由 kubelet 为你分配的。