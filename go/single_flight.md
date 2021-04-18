在给某计算做缓存时遇到了一个问题，描述入下:


![问题描述](./singleflight.png)

针对这个问题，在网友的帮助下，终于搞定。经过调试，[分段锁的方式](./single_flight.go)和[channel同步](single_flight2.go)的方式都是可以的。