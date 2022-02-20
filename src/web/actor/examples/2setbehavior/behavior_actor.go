package _helloworld

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type Person struct {
	Name string
}

type WalkBehavior struct{}

type RunBehavior struct{}

type behaviorActor struct {
	behavior actor.Behavior
}

func (b *behaviorActor) Run(context actor.Context) {
	switch msg := context.Message().(type) {
	case Person:
		context.Respond("dff")
		fmt.Printf("person %v run\n", msg.Name)
	case WalkBehavior:
		b.behavior.Become(b.Walk) // switch behavior
	}
}

func (b *behaviorActor) Walk(context actor.Context) {
	switch msg := context.Message().(type) {
	case Person:
		fmt.Printf("person %v walk\n", msg.Name)
	case RunBehavior:
		b.behavior.Become(b.Run) // switch behavior
	}
}

func NewBehaviorActor() actor.Actor {
	act := &behaviorActor{
		behavior: actor.NewBehavior(),
	}
	act.behavior.Become(act.Run)
	return act
}

func (b *behaviorActor) Receive(ctx actor.Context) {
	b.behavior.Receive(ctx)
}
