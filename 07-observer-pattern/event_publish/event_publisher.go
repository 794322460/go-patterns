package event_publish

import "fmt"

// 事件发送者
type Publisher interface {
	Publish(source EventSource)
}

// 一般一个应用只有一个
type ApplicationPublisher struct {
	dispatcher *EventDispatcher
}

func NewApplicationPublisher(eventDispatcher *EventDispatcher) Publisher {
	fmt.Println("事件发送者：初始化... ")
	return &ApplicationPublisher{dispatcher: eventDispatcher}
}

func (publisher *ApplicationPublisher) Publish(source EventSource) {
	fmt.Println("事件发送者：准备发送事件... ", source)
	publisher.dispatcher.Dispatch(source)
}
