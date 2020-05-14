## 一桩由错误的路径-文件名导致的血案

昨天看到一个叫[rich](https://github.com/willmcgugan/rich)的项目很炫酷，就想尝试一下。

官方有一个使用示例是这样的:
```python
from rich.console import Console
console = Console()

test_data = [
    {"jsonrpc": "2.0", "method": "sum", "params": [None, 1, 2, 4, False, True], "id": "1",},
    {"jsonrpc": "2.0", "method": "notify_hello", "params": [7]},
    {"jsonrpc": "2.0", "method": "subtract", "params": [42, 23], "id": "2"},
]

def test_log():
    enabled = False
    context = {
        "foo": "bar",
    }
    movies = ["Deadpool", "Rise of the Skywalker"]
    console.log("Hello from", console, "!")
    console.log(test_data, log_locals=True)


test_log()

```

把这段代码拷贝过去，命名为 `rich.py`, 然后再在终端执行
```bash
python rich.py
```
错误就在这个时候发生了:
```bash
Traceback (most recent call last):
  File "rich.py", line 1, in <module>
    from rich.console import Console
  File "/Users/xxx/Documents/python/rich.py", line 1, in <module>
    from rich.console import Console
ModuleNotFoundError: No module named 'rich.console'; 'rich' is not a package
```
看到这里，我一直以为是由于机器上`python`版本过多，导致执行
```bash
pip install rich
```
的时候装在了错误的位置，所以找不到。

然后就一直折腾机器上的 `python` 版本:

```bash
➜  ~ ll $(which python3)
lrwxr-xr-x  1 lfn  admin    34B  5 14 02:50 /usr/local/bin/python3 -> ../Cellar/python/3.7.7/bin/python3
➜  ~ python3
Python 3.7.7 (default, May 14 2020, 02:16:58)
[Clang 9.0.0 (clang-900.0.39.2)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> import sys
>>> sys.exc
sys.exc_info(    sys.excepthook(
>>> sys.ex
sys.exc_info(    sys.excepthook(  sys.exec_prefix  sys.executable   sys.exit(
>>> sys.executable
'/usr/local/opt/python/bin/python3.7'
>>>

➜  ~ ll /usr/local/opt/python/bin/python3.7
lrwxr-xr-x  1 lfn  admin    57B  5 14 02:18 /usr/local/opt/python/bin/python3.7 -> ../Frameworks/Python.framework/Versions/3.7/bin/python3.7
```
可见机器上Python版本非常混乱，一时也搞不清楚 `pip install`会安装到哪里，对应哪个 python.

无奈之下，还是想到了安装Python版本管理工具： `anaconda`.

装好之后，进入隔离环境，本来以为这次应该可以了，但是同样的报错还是出现了:
```bash
(mikasa) ➜  ~ python --version
Python 3.8.2
(mikasa) ➜  ~ pip show rich
Name: rich
Version: 1.1.1
Summary: Render rich text, tables, progress bars, syntax highlighting, markdown and more to the terminal
Home-page: https://github.com/willmcgugan/rich
Author: Will McGugan
Author-email: willmcgugan@gmail.com
License: MIT
Location: /Users/lfn/anaconda3/envs/mikasa/lib/python3.8/site-packages
Requires: pygments, commonmark, typing-extensions, pprintpp, colorama
Required-by:


(mikasa) ➜  python rich.py
Traceback (most recent call last):
  File "rich.py", line 1, in <module>
    from rich.console import Console
  File "/Users/lfn/Documents/python/rich.py", line 1, in <module>
    from rich.console import Console
ModuleNotFoundError: No module named 'rich.console'; 'rich' is not a package
(mikasa) ➜  python
```

折腾很长时间，这才意识到问题不出在版本管理上，而是我测试用的py文件名与在py文件中导入的包名重合了。

我们知道, python寻找module的顺序是:

```bash
sys.modules --> built-in --> sys.path
```
看一下`sys.path`
```bash
(mikasa) ➜  python
Python 3.8.2 (default, May  6 2020, 02:49:43)
[Clang 4.0.1 (tags/RELEASE_401/final)] :: Anaconda, Inc. on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> import sys
>>> sys.path
['', '/Users/lfn/anaconda3/envs/mikasa/lib/python38.zip', '/Users/lfn/anaconda3/envs/mikasa/lib/python3.8', '/Users/lfn/anaconda3/envs/mikasa/lib/python3.8/lib-dynload', '/Users/lfn/anaconda3/envs/mikasa/lib/python3.8/site-packages']
>>>
```
可以看到，第一个路径是: `''` , 也就是当下路径。问题到这里就清楚了：
我在rich.py中，第一行就是下面的包导入语句:
```python
from rich.console import Console
```
Python解释器在寻找rich的时候，当前路径会优于 `'/Users/lfn/anaconda3/envs/mikasa/lib/python3.8/site-packages'` 被查找，因为我自己的 `rich.py`只是一个孤立的脚本，并不是一个包(`pachage`), 所以才会报下面的错误:
```bash
ModuleNotFoundError: No module named 'rich.console'; 'rich' is not a package
```

## 小结
给py文件起名的时候要注意，不要与文件中导入的包重名。