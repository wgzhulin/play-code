package _helloworld

import (
	"strconv"
	"testing"

	"github.com/play-code/src/web/actor/examples/internal/stub"
)

func TestBehaviorActor_Receive(t *testing.T) {
	stub := stub.NewStub(NewBehaviorActor)
	defer stub.Finish()

	stub.Msg.Send(Person{Name: "andy"})
	stub.Msg.Send(Person{Name: "jack"})
}

func TestBehaviorActor_ReceiveMultiMsg(t *testing.T) {
	stub := stub.NewStub(NewBehaviorActor)
	defer stub.Finish()

	for i := 0; i < 10; i++ {
		if i == 4 {
			stub.Msg.Send(WalkBehavior{})
		}
		stub.Msg.Send(Person{Name: strconv.Itoa(i)})
	}
}

func TestBehaviorActor_ReceiveSwitchBehavior(t *testing.T) {
	stub := stub.NewStub(NewBehaviorActor)
	defer stub.Finish()
	// default run
	stub.Msg.Send(Person{Name: "andy"})
	// switch behavior walk
	stub.Msg.Send(WalkBehavior{})
	// walk
	stub.Msg.Send(Person{Name: "andy"})
	// switch behavior run
	stub.Msg.Send(RunBehavior{})
	// run
	stub.Msg.Send(Person{Name: "andy"})
	// the output
	// run
	// walk
	// run
}
