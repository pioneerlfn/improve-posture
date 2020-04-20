## http.Request


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