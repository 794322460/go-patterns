package main

import (
	"fmt"
	"henu.yyq/go-patterns/07-observer-pattern/event_publish"
	"reflect"
)

// 事件源A
type EventSourceA struct {
	Id   int
	Name string
}

// 监听器A
type EventListenerA struct{}

func NewEventListenerA() event_publish.EventListener {
	return &EventListenerA{}
}

func (l *EventListenerA) EventType() string {
	return reflect.TypeOf(EventSourceA{}).String()
}

func (l *EventListenerA) OnEventAware(source event_publish.EventSource) {
	fmt.Println("消费者A：接收到消息... ", source)
}

// 事件源B
type EventSourceB struct {
	Id   int
	Name string
}

// 监听器A
type EventListenerB struct{}

func NewEventListenerB() event_publish.EventListener {
	return &EventListenerB{}
}

func (l *EventListenerB) EventType() string {
	return reflect.TypeOf(EventSourceB{}).String()
}

func (l *EventListenerB) OnEventAware(source event_publish.EventSource) {
	fmt.Println("消费者B：接收到消息... ", source)
}

func main() {
	// 初始化事件分发器
	var dispatcher = event_publish.NewEventDispatcher()

	// 初始化事件发送器
	var publisher = event_publish.NewApplicationPublisher(dispatcher)

	// 初始化监听器
	var listenerA = NewEventListenerA()
	var listenerB = NewEventListenerB()

	// 注册事件消费者
	dispatcher.RegisterListener(listenerA)
	dispatcher.RegisterListener(listenerB)

	// 发送事件
	var sourceA = EventSourceA{1001, "张三"}
	var sourceB = EventSourceB{1002, "李四"}
	publisher.Publish(sourceA)
	publisher.Publish(sourceB)

}
