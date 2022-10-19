package publish

import (
	"testing"
)

type Test1 struct{}

func NewTest1() Test1                         { return Test1{} }
func (Test1) Topic() Topic                    { return "test1" }
func (a Test1) Handler(arg interface{}) error { return nil }

type Test2 struct{}

func NewTest2() Test2                         { return Test2{} }
func (Test2) Topic() Topic                    { return "test2" }
func (a Test2) Handler(arg interface{}) error { return nil }

type Test3 struct{}

func NewTest3() Test3                         { return Test3{} }
func (Test3) Topic() Topic                    { return "" }
func (a Test3) Handler(arg interface{}) error { return nil }

func Test_AddSubscriber(t *testing.T) {

	bus := NewEventBus()

	test1 := NewTest1()

	// 添加订阅者
	err := bus.AddSubscriber(test1)
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	// 获取事件
	has := bus.TopicIsExist("test1")
	if !has {
		t.Fail()
	}

	// 添加订阅者
	test2 := NewTest2()
	err = bus.AddSubscriber(test2)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	// 获取事件
	has = bus.TopicIsExist("test2")
	if !has {
		t.Fail()
	}

	// 添加订阅者
	err = bus.AddSubscriber(test1)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	// 获取事件
	has = bus.TopicIsExist("test3")
	if has {
		t.Fail()
	}

	err = bus.AddSubscriber(NewTest3())
	if err.Error() != "topic is null" {
		t.Fail()
	}

}
