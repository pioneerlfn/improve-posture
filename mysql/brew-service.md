## 问题
`brew insatll mysql`之后，mysql并未添加到系统的开机启动及后台常驻项中。

## 解决办法

1. 安装`brew service`命令

    ```bash
    brew tap gapple/services
    ```

2. 启动mysql

    ```bash
    brew service start mysql
    ==> Successfully started `mysql` (label: homebrew.mxcl.mysql)
    ```