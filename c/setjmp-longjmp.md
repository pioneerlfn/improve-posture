这一节我们学习`setjmp`和`longjmp`.

## 概述

`setjmp`和`longjmp`是`nonlocal`的`go`,用于在不同的函数间跳转。

## 要解决的问题
函数调用栈过甚，出现异常或错误的时候，不想层层返回，而是直接回到某个最初的逻辑点，比如main函数中。

## 函数签名
```c
#include <setjmp.h>
int setjmp(jmp_buf env);
> Returns: 0 if called directly, nonzero if returning from a call to longjmp 

void longjmp(jmp_buf env, int val);

```

- setjmp

    直接调用`setjmp`返回0，从调用longjmp返回，返回值是`longjmp`函数中的入参`val`.

- longjmp
    返回到`setjmp`,并且`val`成为`setjmp`返回值。


## 使用示例

```c
#include <setjmp.h>
#include <stdio.h>

static jmp_buf env;

double divide(double to, double by)
{
    if(by == 0)
    {
        longjmp(env, 1);
    }
    return to / by;
}

void test_divide()
{
    divide(2, 0);
    printf("done\n");
}

int main()
{
    if (setjmp(env) == 0)
    {
        test_divide();
    }
    else
    {
        printf("Cannot / 0\n");
        return -1;
    }
    return 0;
}

```


## 注意事项

从`longjmp`返回的时候，`setjmp`所在函数的局部变量是否能保持调用`longjmp`之前的状态是不确定的，如果想要保持这些局部变量仍然保持调用`longjmp`之前的状态，就要用`volatile`修饰。

