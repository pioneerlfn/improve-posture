## 配置文件

### configmap
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: dragonfly-dfdaemon
  labels:
    app: dragonfly
    release: dragonfly
    component: dfdaemon
data:
  dfget.yaml: |-
    aliveTime: 0s
    gcInterval: 1m0s
    keepStorage: false
    workHome: /home/xiaoju/dragonfly-dfdaemon/work
    cacheDir: /home/xiaoju/dragonfly-dfdaemon/cache
    dataDir: /home/xiaoju/dragonfly-dfdaemon/data
    logDir: /home/xiaoju/dragonfly-dfdaemon/logs
    console: false
    verbose: false
    scheduler:
      manager:
        enable: true
        netAddrs:
        - type: tcp
          addr: 10.86.76.24:65003
        refreshInterval: 5m
        seedPeer:
          clusterID: 1
          enable: true
          keepAlive:
            interval: 5s
          type: super
      scheduleTimeout: 30s
      disableAutoBackSource: false
    host:
      idc: ""
      listenIP: 0.0.0.0
      location: ""
      netTopology: ""
      securityDomain: ""
    download:
      calculateDigest: true
      downloadGRPC:
        security:
          insecure: true
          tlsVerify: true
        unixListen:
          socket: /tmp/dfdamon.sock
      peerGRPC:
        security:
          insecure: true
        tcpListen:
          listen: 0.0.0.0
          port: 65000
      perPeerRateLimit: 1024Mi
      totalRateLimit: 2048Mi
    upload:
      rateLimit: 2048Mi
      security:
        insecure: true
        tlsVerify: false
      tcpListen:
        listen: 0.0.0.0
        port: 65002
    storage:
      diskGCThresholdPercent: 90
      multiplex: true
      strategy: io.d7y.storage.v2.simple
      taskExpireTime: 6h
    proxy:
      defaultFilter: Expires&Signature&ns
      tcpListen:
        namespace:
        listen: 0.0.0.0
        port:
      security:
        insecure: true
        tlsVerify: false
      registryMirror:
        dynamic: true
        insecure: false
        url: https://registry-v4.intra.xiaojukeji.com
      proxies:
        - regx: blobs/sha256.*
    objectStorage:
      enable: false
      filter: Expires&Signature&ns
      maxReplicas: 3
      security:
        insecure: true
        tlsVerify: true
      tcpListen:
        listen: 0.0.0.0
        port: 65004
    proxies:
    # 代理镜像 blobs 信息
    - regx: blobs/sha256.*
    # 访问 some-registry 的时候，转换成 https 协议
    - regx: registry-v4.intra.xiaojukeji.com
      useHTTPS: true
    # 直接透传流量，不走蜻蜓
    - regx: no-proxy-reg
      direct: true
    # 转发流量到指定地址
    - regx: some-registry
      redirect: another-registry
    # the same with url rewrite like apache ProxyPass directive
    - regx: ^http://some-registry/(.*)
      redirect: http://another-registry/$1

```

### daemonset

```yaml

apiVersion: v1
items:
- apiVersion: extensions/v1beta1
  kind: DaemonSet
  metadata:
    annotations:
    labels:
      app: dragonfly
      component: dfdaemon
      release: dragonfly
    name: dragonfly-dfdaemon
  spec:
    revisionHistoryLimit: 10
    selector:
      matchLabels:
        app: dragonfly
        component: dfdaemon
        release: dragonfly
    template:
      metadata:
        labels:
          app: dragonfly
          component: dfdaemon
          release: dragonfly
      spec:
        hostNetwork: true
        containers:
        - image: dragonflyoss/dfdaemon:v2.0.4
          imagePullPolicy: IfNotPresent
          livenessProbe:
            exec:
              command:
              - /bin/grpc_health_probe
              - -addr=0.0.0.0:65000
            failureThreshold: 3
            initialDelaySeconds: 15
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 1
          name: dfdaemon
          ports:
          - containerPort: 65001
            protocol: TCP
          readinessProbe:
            exec:
              command:
              - /bin/grpc_health_probe
              - -addr=0.0.0.0:65000
            failureThreshold: 3
            initialDelaySeconds: 5
            periodSeconds: 10
            successThreshold: 1
            timeoutSeconds: 1
          resources:
            limits:
              cpu: "2"
              memory: 2Gi
            requests:
              cpu: "0"
              memory: "0"
          securityContext:
            capabilities:
              add:
              - SYS_ADMIN
            procMount: Default
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /etc/dragonfly
            name: config
          - mountPath: /var/log/dragonfly/daemon
            name: logs
          - mountPath: /run/dragonfly
            name: run
          - mountPath: /var/lib/dragonfly
            name: data
        dnsPolicy: ClusterFirst
        initContainers:
        - command:
          - /bin/sh
          - -cx
          - |-
            if [ ! -e "/run/dragonfly/net" ]; then
              touch /run/dragonfly/net
            fi
            i1=$(stat -L -c %i /host/ns/net)
            i2=$(stat -L -c %i /run/dragonfly/net)
            if [ "$i1" != "$i2" ]; then
              /bin/mount -o bind /host/ns/net /run/dragonfly/net
            fi
          image: dragonflyoss/dfdaemon:v2.0.4
          imagePullPolicy: IfNotPresent
          name: mount-netns
          resources:
            limits:
              cpu: "2"
              memory: 2Gi
            requests:
              cpu: "0"
              memory: "0"
          securityContext:
            privileged: true
            procMount: Default
          terminationMessagePath: /dev/termination-log
          terminationMessagePolicy: File
          volumeMounts:
          - mountPath: /host/ns
            name: hostns
          - mountPath: /run/dragonfly
            mountPropagation: Bidirectional
            name: run
        restartPolicy: Always
        schedulerName: default-scheduler
        securityContext: {}
        terminationGracePeriodSeconds: 30
        volumes:
        - configMap:
            defaultMode: 420
            name: dragonfly-dfdaemon
          name: config
        - hostPath:
            path: /proc/1/ns
            type: ""
          name: hostns
        - hostPath:
            path: /run/dragonfly
            type: DirectoryOrCreate
          name: run
        - emptyDir: {}
          name: data
        - emptyDir: {}
          name: logs
    updateStrategy:
      rollingUpdate:
        maxUnavailable: 1
      type: RollingUpdate
kind: List
metadata:
  resourceVersion: ""
  selfLink: ""

```

### /etc/docker/daemon.json
```json

{
  "registry-mirrors": ["http://127.0.0.1:65002"],
  "live-restore":true,
  "debug":true
}


```

### /etc/hosts
```shell
#rename@(2019-06-11 11:02:07) 127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
::1         localhost localhost.localdomain localhost6 localhost6.localdomain6
127.0.0.1 localhost localhost4 localhost4.localdomain4 #rename@(2019-06-11 11:02:07)
74.125.204.82 k8s.gcr.io
152.199.39.108 get.helm.sh
127.0.0.1 registry-v4.intra.xiaojukeji.com

```

## 拉镜像
```shell
docker pull registry-v4.intra.xiaojukeji.com/didibuild/lfn-debug.py03-pre-v.lfn-debug.deploy.op.didi.com.centos72:dd895446
```

## 验证

```shell
docker exec -it b3908cc08927 sh
cd /home/xiaoju/dragonfly-dfdaemon/logs/daemon/
grep 'peer task done' core.log
```


