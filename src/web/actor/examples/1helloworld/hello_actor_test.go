package _helloworld

import (
	"testing"

	"github.com/play-code/src/web/actor/examples/internal/stub"
)

func TestHelloActor_Receive(t *testing.T) {
	stub := stub.NewStub(NewHelloActor)
	defer stub.Finish()

	stub.Msg.Send(Hello{Name: "andy"})
	stub.Msg.Send(Hello{Name: "jack"})
}
