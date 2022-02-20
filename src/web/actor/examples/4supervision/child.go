package _supervision

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/play-code/src/web/actor/examples/internal/common"
)

type Child struct{}

func NewChild() actor.Actor {
	return &Child{}
}
func (c *Child) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case common.Hello:
		fmt.Println("child handle, hello", msg.Val)
		// child failure
		panic("system err")
	case *actor.Started:
		// do init, like read cfg or load data for database
		fmt.Println("child started")
	case *actor.Stopping:
		fmt.Println("child stopping")
	case *actor.Stopped:
		fmt.Println("child stopped")
	case *actor.Restarting:
		fmt.Println("child restarting")
	default:
		fmt.Printf("unknown commonmsg %T\n", msg)
	}
}
