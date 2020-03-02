# 使用time包应该注意的时区问题

## Unix时间戳

> The unix time stamp is a way to track time as a running total of seconds. This count starts at the Unix Epoch on January 1st, 1970 at UTC. Therefore, the unix time stamp is merely **the number of seconds between a particular date and the Unix Epoch. It should also be pointed out that this point in time technically does not change no matter where you are located on the globe.**

可以看出`unix 时间戳`是一个绝对值，与使用者所在的时区没有关系。
