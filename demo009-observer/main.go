package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataEvent struct {
	Data  interface{} //我们将底层数据定义为一个接口,这意味着它可以是任何值。
	Topic string      //我们另外将主题定义为结构的成员。您的订阅者可能会收听多个主题。因此，我们传递主题是一个很好的做法，以便订阅者可以区分事件。
}

//定义channel来传递DataEvent
// DataChannel is a channel which can accept an DataEvent
type DataChannel chan DataEvent

// DataChannelSlice is a slice of DataChannels
type DataChannelSlice []DataChannel

// EventBus stores the information about subscribers interested for  a particular topic
type EventBus struct {
	subscribers map[string]DataChannelSlice //订阅者，持有 DataChannelSlices 的 map， topic 是map的key
	rm          sync.RWMutex                //来保护它免受读写并发访问
}

//当有人订阅时，我们可以通过key找到topic，然后将event事件传播到channel，做进一步的处理

//Subscribing to a topic
//For subscribing to a topic, a channel is used. It acts like the callback in the traditional approach.
//The channel will receive data when a publisher publishes data to the topic.
//对于订阅主题，使用channel。它的作用类似于传统方法中的回调。当发布者向主题发布数据时，通道将接收数据。
//我们将订阅者附加到通道切片并锁定结构并在操作后解锁它
func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

//Publishing to a topic
//to publish an event, the publisher needs to provide the topic and the data needed to be broadcasted for the subscribers.
//要发布事件，发布者需要为订阅者提供需要广播的主题和数据
func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		// this is done because the slices refer to same array even though they are passed by value
		// thus we are creating a new slice with our elements thus preserve locking correctly.
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

//In this method, first we check if any subscriber exists for the topic. Then we just simply iterate through channel slice associated with the topic and publish it.
//在此方法中，首先我们检查该主题是否存在任何订阅者。然后我们只是简单地遍历与主题关联的通道切片并发布它。

//First, we need to create an instance of the event bus. In a real scenario, you can export a single EventBus from the package making it act like a singleton.
//首先，我们需要创建一个事件总线的实例。在实际场景中，您可以从包中导出单个 EventBus，使其表现得像一个单例。
var eb = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

//To test the newly created event bus we are going to create a method which publishes to a given topic with random intervals
//为了测试新创建的事件总线，我们将创建一个以随机间隔发布到给定主题的方法
func publisTo(topic string, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic: %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}
func main() {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)
	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)
	go publisTo("topic1", "Hi topic 1")
	go publisTo("topic2", "Welcome to topic 2")
	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)
		}
	}
}
