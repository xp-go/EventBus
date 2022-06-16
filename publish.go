package publish

import (
	"context"
	"fmt"
)

func (p Pub) Publish(topic Topic, args interface{}) error {

	method, has := p.event[topic]
	if !has {
		return fmt.Errorf("topic no exist")
	}

	ctx := context.Background()
	err := method.Handler(ctx, args)
	return err
}
