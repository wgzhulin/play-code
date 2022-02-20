package _lifecycleevents

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type PayMethod struct {
	Val string
}

type PaySystem struct{}

func NewPaySystem() actor.Actor {
	return &PaySystem{}
}

func (p *PaySystem) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		// do init, like read cfg or load data for database
		fmt.Println("started")
	case *actor.Stopping:
		fmt.Println("stopping")
	case *actor.Stopped:
		fmt.Println("stopped")
	case *actor.Restarting:
		fmt.Println("restarting")
	case *PayMethod:
		fmt.Println(msg.Val)
	default:
		fmt.Printf("unknown commonmsg %T", msg)
	}
}
