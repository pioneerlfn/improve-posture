## 问题
怎么在terminal漂亮打印json?


## 解决方法
1. jj

    jj是一个用Go写的小工具,使用示例:
    ```bash
    ➜ ~ echo '{"name":"lfn"}' | jj -p
    {
        "name": "lfn"
    }        
    ```
    jj的输出是彩色的，很漂亮。

2. jq
    ```bash
    ➜ ~ echo '{"name":"lfn"}' | jq .

    {
        "name": "lfn"
    }        
    ```

3. python
    ```bash
    ➜ ~ echo '{"name":"lfn"}'  | python -m json.tool

    {
        "name": "lfn"
    }
    ```