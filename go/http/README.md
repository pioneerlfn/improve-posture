## ResponseWriter接口

`ResponseWriter`接口有三个方法:
  - Write
  - WriteHeader
  - Header

`Write`写body, `WriteHeader`写状态码，`Header`设置header.

1. Write
    
    `Write`函数接收`[]byte`类型的参数，并将其写入`ResponseWriter`的body.

    需要注意的是，如果调用`Write`的时候，`content-type`还未设置，则会使用被写入数据的前512个字节探测并设置实际的type.

2. Header

    调用`Header`设置response的header,如:

    ```go
    w.Header.Set("Location", "https://google.com")
    ```
    将把请求重定向到 [google](https://google.com)


3. WriteHeader

    - 调用`WriteHeader`之后，就无法再写response的头部字段了，不过可以继续调用`write`写body.
    - 如果不调用这个方法，那么会自动写入状态码`200 OK`.


