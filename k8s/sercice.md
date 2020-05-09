## Service基础知识

1. `service` 并不是直接通过 `selector` 来转发流量。

    `selector` 被用来生成 `endpoints`资源。
当有请求 `service` 的流量时，`service` 会从对应的`endpoints` 中选择一个转发流量。
2. `endpoints` 不是 `service` 的属性

    `endpoints` 与 `service` 是独立的对象，他们仅通过相同的名字关联。可以手动创建与 `service` 同名的 `endpoints`, 则两者会被k8s自动关联。

3. 集群内的pod 访问集群外的服务有2中方式:
- 手动创建指向外部IP的endpoints
- 将被服务的 `Type` 声明为 `ExternalName`.

