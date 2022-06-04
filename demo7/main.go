package main

import "fmt"

func main() {
	subject := NewSubject()
	// 观察者
	reader1 := NewCustomer("小明")
	reader2 := NewCustomer("小红")
	reader3 := NewCustomer("小李")
	//订阅subject
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	for i := 1; i <= 10; i++ {
		//被观察者 subject 的 context 发生了变更
		subject.UpdateContext(fmt.Sprintf("更新了%d", i))
		fmt.Println("+++++++++++++++++++++++++++++++++")
	}
}

//主题对象
type Subject struct {
	observers []Observer //观察者们
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

//观察者订阅主题后，追加到
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

//遍历观察者们，通知变更
func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

//观察者
type Observer interface {
	Update(*Subject)
}

//Customer 实现了 Observer
type Customer struct {
	name string
}

func NewCustomer(name string) *Customer {
	return &Customer{
		name: name,
	}
}

//
func (r *Customer) Update(s *Subject) {
	fmt.Printf("%s received %s\n", r.name, s.context)
}
