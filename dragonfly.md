- 使用helm install
- 查看结果：
root@ddcloud-kube-master00-v1.py:~$ kubectl get po -n dragonfly-system
NAME                                 READY     STATUS                  RESTARTS   AGE
dragonfly-dfdaemon-2vqlx             0/1       Pending                 0          3h
dragonfly-dfdaemon-cv8js             0/1       Init:ImagePullBackOff   0          3h
dragonfly-dfdaemon-dw4gf             0/1       Terminating             0          5h
dragonfly-dfdaemon-gfsdm             0/1       Init:ImagePullBackOff   0          3h
dragonfly-dfdaemon-kb7xj             0/1       Terminating             0          5h
dragonfly-dfdaemon-l6nzt             0/1       Error                   5          3h
dragonfly-dfdaemon-nvlsp             0/1       Error                   5          3h
dragonfly-dfdaemon-qswwg             0/1       Init:CrashLoopBackOff   21         3h
dragonfly-manager-58b949cd4c-7hvh6   0/1       Pending                 0          1m
dragonfly-manager-58b949cd4c-g6vr5   0/1       Pending                 0          1m
dragonfly-manager-58b949cd4c-ng4sq   0/1       Pending                 0          1m
dragonfly-manager-6f6d7b577d-dgf58   0/1       CrashLoopBackOff        2          1m
dragonfly-scheduler-0                0/1       Init:0/1                0          1h
dragonfly-seed-peer-0                0/1       Pending                 0          3h

- 查看日志
```shell
kubectl logs dragonfly-manager-6f6d7b577d-dgf58 -n dragonfly-system

panic: fqdn hostname not found: lookup dragonfly-manager-6f6d7b577d-dgf58. on 10.85.131.231:53: no such host

goroutine 1 [running]:
d7y.io/dragonfly/v2/pkg/net/fqdn.fqdnHostname(...)
	/go/src/d7y.io/dragonfly/v2/pkg/net/fqdn/fqdn.go:33
d7y.io/dragonfly/v2/pkg/net/fqdn.init.0()
	/go/src/d7y.io/dragonfly/v2/pkg/net/fqdn/fqdn.go:26 +0x65
  
``` 

登录dragonfly-manager-6f6d7b577d-dgf58所在宿主机
使用镜像，指定entrypoint登录

```shell
docker run -it --entrypoint sh 88fa64e6fef7
/opt/dragonfly # hostname
e1e8c3decaee
/opt/dragonfly # cat /etc/hosts
127.0.0.1	localhost
::1	localhost ip6-localhost ip6-loopback
fe00::0	ip6-localnet
ff00::0	ip6-mcastprefix
ff02::1	ip6-allnodes
ff02::2	ip6-allrouters
172.17.0.2	e1e8c3decaee

```
/opt/dragonfly # hostname
e1e8c3decaee

cd /opt/dragonfly/bin
/opt/dragonfly/bin # ls
manager  server
/opt/dragonfly/bin # ./server

执行.server没有任何输出

  
