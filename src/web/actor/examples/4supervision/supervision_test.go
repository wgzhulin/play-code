package _supervision

import (
	"fmt"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/play-code/src/web/actor/examples/internal/stub"
	"testing"
)

func TestParent_Receive(t *testing.T) {
	stub := stub.NewStub(NewParent)
	defer stub.Finish()

	stub.Msg.Send(crateChild{})
}

func TestParent_ReceiveChildFail(t *testing.T) {
	decider := func(reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		// 监管策略：继续处理，停止子Actor，重启子Actor，升级上报给父Actor的父Actor
		return actor.StopDirective
	}

	supervisor := actor.NewOneForOneStrategy(10, 1000, decider)
	props := actor.
		PropsFromProducer(NewParent).
		WithSupervisor(supervisor)

	stub := stub.NewStubProps(props)
	defer stub.Finish()

	stub.Msg.Send(crateChild{})
}