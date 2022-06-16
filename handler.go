package publish

import (
	"context"
	"fmt"
	"reflect"
)

type Pub struct {
	event map[Topic]Handler
}

func NewPub() Pub {
	return Pub{
		event: make(map[Topic]Handler),
	}
}

func (p Pub) Register(handler Handler) error {

	of := reflect.TypeOf(handler)
	name, ok := of.MethodByName("Topic")
	if !ok {
		return fmt.Errorf("topic function no exist")
	}

	valueOf := reflect.ValueOf(handler)
	results := name.Func.Call([]reflect.Value{valueOf})
	if len(results) != 1 {
		return fmt.Errorf("topic function return args error")
	}

	_, has := of.MethodByName("Handler")
	if !has {
		return fmt.Errorf("handler function no exist")
	}
	p.event[Topic(results[0].String())] = handler

	return nil
}

type Handler interface {
	Topic() Topic
	Handler(ctx context.Context, args interface{}) error
}
