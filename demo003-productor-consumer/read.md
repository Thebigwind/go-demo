生产者消费者模型： 一个生产者，多个消费者

一个生产者生产，多个消费者消费
生产者生产数据，发送到chan； 生产者生产完后，关闭channel ；
消费者从chan中取数据，进行消费逻辑处理；获取到channel关闭的信息后，消费者goroutine主动退出；

 ----------------------

生产者消费者模型： 多个生产者，多个消费者

和一个生产者，多个消费者模型区别，主要是处理多个生产者。

通过 wg.Add 起多个生产者；
当生产者没有数据可发送到channel时，wg.Done(），生产者退出；
当全部的生产者都退出后，wg.Wait()执行结束，此时 close channel.

