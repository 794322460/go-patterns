package event_publish

// 消费者接口
// PS：多了个EventType来判断类型，因为没有Javaの范型支持，Java写法如下：
// public interface EventListener<T extends EventSource> {}
// 可以通过T拿到具体类型
type EventListener interface {
	EventType() string
	OnEventAware(source EventSource)
}
