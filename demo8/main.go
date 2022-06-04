// design_mode/pkg/observer/observer.go
package main

import (
	"fmt"
	"sync"
)

type (
	Event struct {
		Data int64
	}
	//观察者
	Observer interface {
		// allows an event to be "published" to interface implementations
		OnNotify(Event)
	}
	// 被观察者
	Notifier interface {
		Register(Observer)
		Deregister(Observer)
		// notify new events to observers
		Notify(Event)
	}
)

// MyNotifier 实现 Notifier 接口
type MyNotifier struct {
	sync.RWMutex
	observers []Observer
}

func NewMyNotifier() Notifier {
	return &MyNotifier{
		observers: make([]Observer, 0),
	}
}

func (m *MyNotifier) Register(o Observer) {
	m.Lock()
	m.observers = append(m.observers, o)
	m.Unlock()
}

func (m *MyNotifier) Deregister(o Observer) {
	m.RLock()
	pos := -1
	for i := range m.observers {
		if o == m.observers[i] {
			pos = i
			break
		}
	}
	m.RUnlock()
	if pos != -1 {
		m.Lock()
		m.observers = append(m.observers[0:pos], m.observers[pos+1:]...)
		m.Unlock()
	}
}

func (m *MyNotifier) Notify(event Event) {
	wg := sync.WaitGroup{}
	wg.Add(len(m.observers))

	for i := range m.observers {
		go func(i int) {
			defer wg.Done()
			m.RLock()
			m.observers[i].OnNotify(event)
			m.RUnlock()
		}(i)
	}
	wg.Wait()
}

// ObserverA 观察者，实现 Observer 接口
type ObserverA struct {
}

func (o *ObserverA) OnNotify(event Event) {
	fmt.Printf("ObserverA receive event:%v\n", event)
}

// ObserverB 观察者，实现 Observer 接口
type ObserverB struct {
}

func (o *ObserverB) OnNotify(event Event) {
	fmt.Printf("ObserverB receive event:%v\n", event)
}

func main() {
	notify := NewMyNotifier()
	a := &ObserverA{}
	b := &ObserverB{}

	notify.Register(a)
	notify.Register(b)
	notify.Notify(Event{Data: 1})

	notify.Deregister(a)
	notify.Notify(Event{Data: 2})
}
