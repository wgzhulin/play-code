package stub

import (
	"fmt"
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type MsgHandler struct {
	ctx      *actor.RootContext
	receiver *actor.PID
}

func (m *MsgHandler) Send(msg interface{}) {
	m.ctx.Send(m.receiver, msg)
}

// sends a message to a given PID and returns a Future
func (m *MsgHandler) RequestFuture(msg interface{}) {
	future := m.ctx.RequestFuture(m.receiver, msg, time.Second)
	rsp, err := future.Result()
	if err != nil {
		fmt.Printf("RequestFuture %v\n", err)
	} else {
		fmt.Printf("RequestFuture rsp is %v\n", rsp)
	}
}

func (m *MsgHandler) SendStop() {
	m.ctx.Stop(m.receiver)
}
