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

## API返回值

HTTP Code
很多API的HTTP Code没有被用到，这是非常不恰当的。比如很多接口的HTTP Code 永远都是返回200，200只是表示接口没有挂。

所以还是需要列一下常用的HTTP Code：

| HTTP CODE| 语义 | 常规用途 |
| ---       | --- | --- |
| 200 | OK | GET接口并且返回了正确值; <br>PUT接口修改成功后返回 |
| 201 | Created |	POST接口成功返回 |
| 202 | Accepted | POST接口成功返回，**用于异步任务** |
| 204 |	No Content |	DELETE接口成功后返回 |
| 400 | Bad Request |	传递了错误参数 |
| 401 |	Unauthorized |	缺少用户登录信息 |
| 403 |	Forbidden |	用户权限不足 |
| 404 |	Not Found |	没有找到相关信息 |
| 409 |	Conflict |	冲突 |
| 500 | Internal Server Error |	内部错误，通常是bug或者是下游接口错误 |

## Html-form
We know the data is sent to the server through an HTTP POST request and is placed in the body of the request. But how is the data formatted? The HTML form data is always sent as name-value pairs, but how are these name-value pairs formatted in the POST body? It’s important for us to know this because as we receive the POST request from the browser, we need to be able to parse the data and extract the name-value pairs.

The format of the name-value pairs sent through a POST request is specified by the content type of the HTML form. This is defined using the enctype attribute like this:

```html
<form action="/process" method="post" enctype="application/x-www-form-urlencoded">
  <input type="text" name="first_name"/>
  <input type="text" name="last_name"/>
  <input type="submit"/>
</form>
```

The default value for enctype is `application/x-www-form-urlencoded`. Browsers are required to support at least application/x-www-form-urlencoded and multipart/ form-data (*HTML5 also supports a text/plain value*).

If we set enctype to `application/x-www-form-urlencoded`, the browser will encode in the HTML form data a long query string, with the name-value pairs sepa- rated by an ampersand (`&`) and the name separated from the values by an equal sign (`=`). That’s the same as URL encoding, hence the name.
In other words, the HTTP body will look something like this:
```
first_name=sau%20sheong&last_name=chang
```

If you set enctype to `multipart/form-data`, each name-value pair will be converted into a MIME message part, each with its own content type and content disposition. Our form data will now look something like this:
```
------WebKitFormBoundaryMPNjKpeO9cLiocMw 
Content-Disposition: form-data; name="first_name"

sau sheong 
------WebKitFormBoundaryMPNjKpeO9cLiocMw Content-Disposition: form-data; name="last_name"

chang
------WebKitFormBoundaryMPNjKpeO9cLiocMw--

```
When would you use one or the other? If you’re sending simple text data, the URL encoded form is better—it’s simpler and more efficient and less processing is needed. If you’re sending large amounts of data, such as uploading files, the multipart-MIME form is better. You can even specify that you want to do Base64 encoding to send binary data as text.