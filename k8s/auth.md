# 认证相关

## secret卷
挂载路径是: `/var/run/secrets/kubernetes.io/serviceaccount`
每个pod默认挂载secret 卷，下面三个文件，其用途分别为:
- `token` 用来访问apiserver的token
- `ca.crt` 用来认证服务端(apiserver)发来的证书
- `namespace` 表征当前pod所在的ns
> Let’s recap how an app running inside a pod can access the Kubernetes API properly:<br>
>  The app should verify whether the API server’s certificate is signed by the certif- icate authority, whose certificate is in the ca.crt file.<br>
>  The app should authenticate itself by sending the Authorization header with the bearer token from the token file.<br>
>  The namespace file should be used to pass the namespace to the API server when performing CRUD operations on API objects inside the pod’s namespace.<br>




