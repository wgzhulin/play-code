package _helloworld

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type Hello struct {
	Name string
}

type helloActor struct{}

func NewHelloActor() actor.Actor {
	return &helloActor{}
}

func (h *helloActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Name)
	}
}
