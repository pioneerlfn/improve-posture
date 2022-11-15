- [Managing a TLS Certificate for Kubernetes Admission Webhook](https://www.velotio.com/engineering-blog/managing-tls-certificate-for-kubernetes-admission-webhook)
简单来说，webhook自签的证书，只要apiserver认就行. 
apiserver怎么知道呢？
webhook 通过client-go将签发自己证书的ca写入mutationwebhookconfiguration的caBundle里就行。
当然这个ca也是webhook自己生成的，用来签发自己证书的.
