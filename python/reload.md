## 模块热加载
python的模块热加载机制，让我们有能力不重启解释器的情况下改变模块的代码，个人感觉是个比较有意思的功能。

举个例子，假设我们在当前目录下有一个叫`diamond`的模块，可以打印class的继承顺序。

```python
# diamond.py

class A:
    def ping(self):
        print("ping", self)


class B(A):
    def pong(self):
        print("pong", self)

class C(A):
    def pong(self):
        print("PONG", self)

class D(B,C):
    def ping(self):
        super().ping()
        print("post-ping", self)
    
    def pingpong(self):
        self.ping()
        super().ping()
        self.pong()
        super.pong()
        C.pong(self)

def print_mro(cls):    
    print(', '.join(c.__name__ for c in cls.__mro__))

```
然后我们打开解释器：
```shell
(mikasa) ➜  python python
Python 3.8.2 (default, May  6 2020, 02:49:43)
[Clang 4.0.1 (tags/RELEASE_401/final)] :: Anaconda, Inc. on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> import diamond
>>> import tkinter
>>> diamond.print_mro(tkinter.Text)
Text, Widget, BaseWidget, Misc, Pack, Place, Grid, XView, YView, object

```

然后我们给`print_mro`函数中加入一行,函数变成下面这样:
```python
def print_mro(cls):    
    print('diamond.print_mro is called...')
    print(', '.join(c.__name__ for c in cls.__mro__))
```
在解释器中重载`diamond`模块:
```shell
>>> from importlib import reload
>>> reload(diamond)
<module 'diamond' from '/Users/lfn/Documents/python/diamond.py'>
>>> diamond.print_mro(tkinter.Text)
diamond.print_mro is called...
Text, Widget, BaseWidget, Misc, Pack, Place, Grid, XView, YView, object
>>>
```

可以看出来，更改生效了。