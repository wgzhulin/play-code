package _lifecycleevents

import (
	"testing"

	"github.com/play-code/src/web/actor/examples/internal/stub"
)

func TestPaySystem_Receive(t *testing.T) {
	// spawn
	stub := stub.NewStub(NewPaySystem)
	defer stub.Finish()

	stub.Msg.Send(&PayMethod{Val: "weixin"})
}
