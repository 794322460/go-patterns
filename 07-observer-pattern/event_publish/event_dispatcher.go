package event_publish

import (
	"fmt"
	"reflect"
)

type EventDispatcher struct {
	// 做广播用
	broadcastReceiver []EventListener
	// 每个事件类型下注册の消费者（是个list）
	eventListenerMap map[string][]EventListener
}

func NewEventDispatcher() *EventDispatcher {
	fmt.Println("事件分发器：初始化...")
	return &EventDispatcher{
		broadcastReceiver: make([]EventListener, 0, 10),
		eventListenerMap:  make(map[string][]EventListener, 10),
	}
}

func (d *EventDispatcher) AddBroadcastReceiver(receiver EventListener) {
	fmt.Println("事件分发器：添加广播接收者... ", receiver)
	d.broadcastReceiver = append(d.broadcastReceiver, receiver)
}

func (d *EventDispatcher) RegisterListener(listener EventListener) {
	fmt.Println("事件分发器：注册事件接收者... ", listener.EventType(), "-", reflect.TypeOf(listener).Elem().String())
	listeners := d.eventListenerMap[listener.EventType()]
	if listeners == nil {
		listeners = make([]EventListener, 0, 10)
		// d.eventListenerMap[listener.EventType()] = listeners
	}
	listeners = append(listeners, listener)
	// PS：为啥放到这里就好了？   学Javaの我实在有点难以接受
	d.eventListenerMap[listener.EventType()] = listeners
}

func (d *EventDispatcher) Dispatch(source EventSource) {
	fmt.Println("事件分发器：准备分发事件... ", source)
	if source == nil {
		return
	}
	// 获取类型
	typeName := reflect.TypeOf(source).String()
	listeners := d.eventListenerMap[typeName]
	for _, listener := range listeners {
		fmt.Println("事件分发器：开始分发事件... ", listener.EventType(), "-", reflect.TypeOf(listener).Name(), ":", source)
		listener.OnEventAware(source)
	}
}

func (d *EventDispatcher) Broadcast(source EventSource) {
	for _, receiver := range d.broadcastReceiver {
		receiver.OnEventAware(source)
	}
}
