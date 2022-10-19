package publish

import (
	"fmt"
	"reflect"
)

// 事件总线
type Bus interface {
	Publisher
	Subscriber
}

// 订阅者
type Subscriber interface {
	// 添加订阅者
	AddSubscriber(handler Handler) error
	// 查找事件是否存在
	TopicIsExist(topic Topic) bool
}

// 发布者
type Publisher interface {
	// 事件发布
	Publish(topic Topic, args interface{}) error
}

// 订阅者需要实现的接口
type Handler interface {
	Topic() Topic
	Handler(args interface{}) error
}

type Topic string

type EventBus struct {
	event map[Topic][]Handler
}

func NewEventBus() Bus {
	bus := &EventBus{
		event: make(map[Topic][]Handler),
	}
	return bus
}

func (e EventBus) AddSubscriber(handler Handler) error {

	of := reflect.TypeOf(handler)
	name, ok := of.MethodByName("Topic")
	if !ok {
		return fmt.Errorf("Topic function no exist")
	}

	valueOf := reflect.ValueOf(handler)
	results := name.Func.Call([]reflect.Value{valueOf})
	if len(results) != 1 {
		return fmt.Errorf("Topic function return args error")
	}

	_, has := of.MethodByName("Handler")
	if !has {
		return fmt.Errorf("Handler function no exist")
	}
	topic := handler.Topic()
	if topic == "" {
		return fmt.Errorf("topic is null")
	}

	handlers, has := e.event[topic]
	if !has {
		handlers = make([]Handler, 0)
	}
	handlers = append(handlers, handler)

	e.event[topic] = handlers

	return nil
}

func (e EventBus) TopicIsExist(topic Topic) bool {
	_, has := e.event[topic]
	return has
}

func (p EventBus) Publish(topic Topic, args interface{}) error {
	handlers, has := p.event[topic]
	if !has {
		return fmt.Errorf("topic no exist")
	}

	for _, handler := range handlers {
		err := handler.Handler(args)
		if err != nil {
			return fmt.Errorf("handler exist error %s", topic)
		}
	}
	return nil
}
