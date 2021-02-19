- `kubectl exec kubia-7nog1 -- curl -s http://10.111.249.153`中的 `--` 用来标志`kubectl` 命令选项的结束。所以 `--` 之后的命令会在 `pod` 内部执行;
- 如果想让同一个 `client` 的请求总是被转发到特定的 `pod`, 可以用 `sessionAffinity: ClientIP`:

	```yaml
	apiVersion: v1
	kind: Service
	spec:
		sessionAffinity: ClientIP
		...
	```
	`kubernetes` 只支持下面两种 `sessionAffinity`:
	- None
	- ClientIP
	
	`service` 工作在4层(`TCP/UDP`)
	
- `expose` 多端口	

	```yaml
	apiVersion: v1
	kind: Service
	metadata: null
	name: kubia
	spec:
	  ports:
	    - name: http
	      port: 80
	      targetPort: 8080
	    - name: https
	      port: 443
	      targetPort: 8443
	  selector:
	    app: kubia
	```
	当使用多端口的时候，必须为每一个端口起一个名字
	

### service代理外部服务	

当 `service` 中不包含 `selector` 的时候，可以通过创建**同名** `Endpoints` 的方式手动指定service代理的实际服务，通过这种方式，可以让 `cluster`内部的 `pod` 访问 `cluster` 外部的服务。

























	
	
