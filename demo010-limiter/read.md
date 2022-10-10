限流器 Limiter 

限流器的实现方法有很多种，例如 Token Bucket、滑动窗口法、Leaky Bucket等。

在 Golang 库中官方给我们提供了限流器的实现golang.org/x/time/rate，它是基于令牌桶算法（Token Bucket）设计实现的。


https://github.com/golang/time