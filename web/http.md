## HTTP动词

| Method | Meaning | Safe | Idempotent |
|--- | --- | --- | --- |
| GET | Tells the server to return the specified resource. | YES | YES |
| HEAD | The same as GET except that the server must not return a message body. This method is often used to get the response headers without carrying the weight of the rest of the message body over the network. | YES | YES |
| OPTIONS | Tells the server to return a list of HTTP methods that the server sup-ports. | YES | YES |
| TRACE | Tells the server to return the request. This way, the client can see what the intermediate servers did to the request. | YES | YES |
| POST | Tells the server that the data in the message body should be passed to the resource identified by the URI. What the server does with the message body is up to the server. | **NO** | **NO** |
| PUT | Tells the server that the data in the message body should be the resource at the given URI. If data already exists at the resource identified by the URI, that data is replaced. Otherwise, a new resource is created at the place where the URI is. | NO | YES |
| DELETE | Tells the server to remove the resource identified by the URI. | NO | YES |
| CONNECT | Tells the server to set up a network connection with the client. This method is used mostly for setting up SSL tunneling (to enable HTTPS). |  |  |
| PATCH | Tells the server that the data in the message body modifies the resource identified by the URI. | | | 

## Tips

HTTP头部字段是 key-value 的形式，key 和 value 之间用“:”分隔，最后用 CRLF 换行表示字段结束。比如在“Host: 127.0.0.1”这一行里 key 就是“Host”，value 就是“127.0.0.1”。

HTTP 头字段非常灵活，不仅可以使用标准里的 Host、Connection 等已有头，也可以任意添加自定义头，这就给 HTTP 协议带来了无限的扩展可能。

不过使用头字段需要注意下面几点：
- 字段名不区分大小写，例如“Host”也可以写成“host”，但首字母大写的可读性更好；
- 字段名里不允许出现空格，可以使用连字符“-”，但不能使用下划线“_”。例如，“test-name”是合法的字段名，而“test name”“test_name”是不正确的字段名；- 字段名后面必须紧接着“:”，不能有空格，而“:”后的字段值前可以有多个空格；
- 字段的顺序是没有意义的，可以任意排列不影响语义；字段原则上不能重复，除非这个字段本身的语义允许，例如 Set-Cookie。

- HTTP request headers are mostly optional. The only mandatory header in HTTP 1.1 is the Host header field. But if the message has a message body (which is optional, depending on the method), you’ll need to have either the Content-Length or the Transfer-Encoding header fields

