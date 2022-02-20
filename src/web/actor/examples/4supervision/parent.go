package _supervision

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/play-code/src/web/actor/examples/internal/common"
)

// crate child actor
type crateChild struct{}

type Parent struct{}

func NewParent() actor.Actor {
	return &Parent{}
}
func (c *Parent) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case common.Hello:
		fmt.Println("hello", msg.Val)
	case crateChild:
		props := actor.PropsFromProducer(NewChild)
		pid := ctx.Spawn(props)
		ctx.Send(pid, common.Hello{Val: "create child"})
	}
}
