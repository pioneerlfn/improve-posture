## or-done-channel 模式

当我们工作与一个比较复杂的系统时，有时候我们的代码会从其他组件提供的channel中读数据。

由于我们无法掌控他人的channel关闭时机，当我们自己的goroutine被取消时，我们无法确认我们自己goroutine中读到的channel是否被关闭。

比如当我们采用下面的模式从别人channel读数据时，有可能会被阻塞,导致我们的goroutine无法及时退出。

```Go
for val := range othersChan {
    // do sth
}
```

为了解决这个问题，我们可以采用下面被称为`or-done-channel`的模式.

```Go
func orDone(done, c <-chan interface{}) <-chan interface{} { 
	valStream := make(chan interface{})
	go func() {
		defer close(valStream) 
		for {
			select { 
            // 1.即使别人的channle不关闭(情况2), 我们的goroutine也可以退出
            case <-done:
				return
			case v, ok := <-c:
				if ok == false { 
                    // 2.如果别人的channel关闭了，则返回
					return
				}
                // 如果别人的channel没关闭                
                select {
				case valStream <- v: 
                // 3. 当前goroutine被要求退出,为了保证不阻塞在发送上                
                case <-done:
				}
			} 
		}
	}()
	return valStream 
}

```

然后我们便可以像下面这样读，而无需担心阻塞在别人的channel上:
```Go
for val := range orDone(done, othersChan) {
    // do sth.
}
```




