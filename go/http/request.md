## http.Request


### URL query
一般的请求格式为:

```
scheme://[userinfo@]host/path[?query][#fragment]
```

在`scheme`之后不紧跟`/`的URL会被翻译为:

```
scheme:opaque[?query][#fragment]
```

`Request`有一个字段是`URL`

```Go
type Request struct {
    ...
    
    URL *url.URL
    
    ...
}
```

```Go
type URL struct {
    Scheme     string
	Opaque     string    // encoded opaque data
	User       *Userinfo // username and password information
	Host       string    // host or host:port
	Path       string    // path (relative paths may omit leading slash)
	RawPath    string    // encoded path hint (see EscapedPath method)
	ForceQuery bool      // append a query ('?') even if RawQuery is empty
	RawQuery   string    // encoded query values, without '?'
	Fragment   string    // fragment for references, without '#'
}

```

可以从`http.Request.URL`的`RawQuey`字段获取到请求的query部分。举个例子🌰:

如果我们有一个请求的URL是:
```
http://www.example .com/post?id=123&thread_id=456
```
那么，`RawQuery`中的内容即为:

```
id=123&thread_id=456
```
 
我们需要解析`RawQuery`来获得查询键值对。

### form（POST发送)
请求的表单数据，其编码方式是由`form`的`enctype`决定的，比如下面这个表单，就定义了name-value的编码形式是`x-www-form-urlencoded`,这也会表单的默认编码方式:

```html
<form action="/process" method="post" enctype="application/x-www-form-urlencoded">
  <input type="text" name="first_name"/>
  <input type="text" name="last_name"/>
  <input type="submit"/>
</form>
```

浏览器被要求至少支持以下表单编码方式：
- x-www-form-urlencoded
- multipart/form-data
- text/plain(html5要求)

对于`x-www-form-urlencoded`来说，编码后的内容是一个比较长的`qeury string`,比如：

```
first_name=sau%20sheong&last_name=chang
```
这和`URL查询`的编码一样，因此叫这个名字。

如果是简单的数据，使用`x-www-urlencoded`就可以了。如果你需要发送大量的数据，比如上传文件等，那`multipart-form-data`这种形式更好。

### form(GET发送)
`GET`同样也可以发送表单，这一点没怎么注意到。

`GET`请求是没有`request body`的，因此可以将表单以`URL`的形式发送，只需要将`method`设置为`get`. 比如:

```html
<form action="/process" method="get"> 
          <input type="text" name="first_name"/>
          <input type="text" name="last_name"/>
          <input type="submit"/>
</form>
```


## Form解析
虽然我们可以手动解析请求的`url`和`body`,但这通常是没必要的。我们可以利用标准库提供的一系列辅助函数完成解析Form的任务。

通常的步骤是:
1. 调用Call `ParseForm` 或 `ParseMultipartForm` 解析请求
2. 访问`Form`, `PostForm`, `MultipartForm`获取数据。


假设我们发送了下面这个请求:

```html
<form action=http://127.0.0.1:8080/process?hello=world&thread=123 method="post" enctype="application/x-www-form-urlencoded">
	<input type="text" name="hello" value="sau sheong"/>
	<input type="text" name="post" value="456"/>
	<input type="
</form>
```

### 获取form表单和URL查询参数
调用`Request.ParseForm`解析之后，我们便可以从`Form`字段获得以下内容:

```
map[thread:[123] hello:[sau sheong world] post:[456]]
```

可以看出:
- `Form`中是一个`map[string][]string`,其中`hello`的val是`[]string{"sau sheong", "world"}`, 
- "sau sheong"来自form表单
- "world"来自url查询参数

>**注意，对同一个key, 来自form表单的val总是在来自url查询参数中的val之前。**

### 只获取form表单数据

从上面👆内容可知，调用`ParseForm`方法解析原始请求之后，在`Form`字段中获取的数据，既有来自form表单的，也有来自url查询参数的。

如果我们我们只想获取form表达的数据，该怎么办呢？

**访问`PostForm`字段。**

还是上面的例子，访问`PostForm`得到的内容是:
```
map[post:[456] hello:[sau sheong]]
```

### MultipartForm

在`Form`字段，只能获取到`x-www-urlencoded`编码的数据，如果想获取以`multipart/form-data`编码的表单内容，那就需要:

1. 调用`ParseMultiPartForm`
2. 访问`MultipartForm`

`MultipartForm`中不包含`url查询参数`, 只有form表单数据：
- 有2个字典
- 第一个字典内容是写在html表单的key-value
- 第二个字典是关于文件的。如果没有上传文件，则为空字典



### 捷径方法: FormValue()

常规来讲，我们需要先调用`ParseForm`，然后再访问`Form`内容。
但是我们使用`FormValue`方法就不一样了:
- 如果需要，`FormValue`会自动调用`ParseForm`
- `FormParse(key)`只会返回第一个`key`对应切片的第一个元素

### PostFormValue()

`PostFormValue`与`FormValue`的原理类似，只不过不会去解析`url查询参数`.

### 一个坑
`FormValue()`和`PostFormValue()`都会调用`ParseMultiPartForm`来解析数据。如果客户端的表单编码形式是`multipart/form-data`,那调用这俩货之后，无法在`Form`或者`PostForm`字段中拿到请求数据，反而可以在`MultiPartForm`中拿到。
> 注意: `ParseMultiPartForm()`在`r.Form==nil`的时候，会调用`ParseForm()`.





## 推荐阅读
[Go Web Programming](https://github.com/KeKe-Li/book/blob/master/Go/go-web-programming.pdf)