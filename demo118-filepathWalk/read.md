Golang回调函数

Golang回调函数实例二则

#定义
回调函数就是一个通过函数指针调用的函数。如果你把函数的指针（地址）作为参数传递给另一个函数，当这个指针被用来调用其所指向的函数时，我们就说这是回调函数。
回调函数不是由该函数的实现方直接调用，而是在特定的事件或条件发生时由另外的一方调用的，用于对该事件或条件进行响应。

#意义
回调函数是用户实现异步的一种方式：把处理函数注册为一个路由的回调函数，当有请求后自动调用回调函数； 这样主程序的执行，就不受到请求的影响，实现了异步。
当然这里的异步机制由epoll实现，不能算严格意义上的异步。

#机制
##定义一个回调函数
提供函数实现的一方在初始化的时候，将回调函数的函数指针注册给调用者