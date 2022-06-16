package publish

import (
	"context"
	"fmt"
	"testing"
)

type AddGold struct {
	Name string
}

func (AddGold) Topic() Topic {
	return "ssss"
}

func (a AddGold) Handler(ctx context.Context, arg interface{}) error {
	fmt.Println("sss")
	fmt.Println(a.Name)
	return nil
}

func Test_Demo(t *testing.T) {
	pub := NewPub()
	err := pub.Register(AddGold{Name: "cxp"})
	pub.Publish("ssss", nil)
	fmt.Println(err)
}
