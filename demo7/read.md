`观察者模式，事件机制使用

**观察者模式**
定义
观察者模式(Observer Design Pattern)定义了一种一对多的依赖关系，让多个观察者对象同时监听一个主题对象。这个主题对象在状态发生变化的时，会通知所有的观察者对象，使他们能够更新自己。
定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新

**适用场景**
1、当一个对象状态的改变需要改变其他对象，或实际对象是事先未知的或动态变化的，可使用观察者模式；
2、当应用中的一些对象必须观察其他对象时，可使用该模式。但仅能在有限时间内或特定情况下使用。

优点
1、降低了目标与观察者之间的耦合关系，两者之间是抽象耦合关系。
2、目标与观察者之间建立了一套触发机制。

缺点
`1、目标与观察者之间的依赖关系并没有完全解除，而且有可能出现循环引用。
2、当观察者对象很多时，通知的发布会花费很多时间，影响程序的效率。
`

不同场景的实现方式
针对应用场景有下面四种实现方式

    1、同步阻塞的实现方式；
    2、异步非阻塞的实现方式；
    3、进程内的实现方式；
    4、跨进程的实现方式。

栗如：可以基于消息队列实现

被观察者直接通知到观察者这种场景就是同步阻塞的实现方式。


gitHub: https://github.com/asaskevich/EventBus
文档
https://pkg.go.dev/github.com/asaskevich/eventbus#NewServer


`