package stub

import (
	"time"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type Stub struct {
	root *actor.RootContext
	pid  *actor.PID
	Msg  *MsgHandler
}

func NewStub(newActor func() actor.Actor) *Stub {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(newActor)
	// default cfg
	pid := system.Root.Spawn(props)
	return &Stub{
		root: system.Root,
		pid:  pid,
		Msg:  &MsgHandler{ctx: system.Root, receiver: pid},
	}
}

func NewStubReceiveFunc(receiveFunc func(ctx actor.Context)) *Stub {
	system := actor.NewActorSystem()
	props := actor.PropsFromFunc(receiveFunc)
	// default cfg
	pid := system.Root.Spawn(props)
	return &Stub{
		root: system.Root,
		pid:  pid,
		Msg:  &MsgHandler{ctx: system.Root, receiver: pid},
	}
}

func NewStubProps(props *actor.Props) *Stub {
	system := actor.NewActorSystem()
	// default cfg
	pid := system.Root.Spawn(props)
	return &Stub{
		root: system.Root,
		pid:  pid,
		Msg:  &MsgHandler{ctx: system.Root, receiver: pid},
	}
}

func (s *Stub) Finish() {
	// wait print
	time.Sleep(time.Millisecond * 10)
	s.root.StopFuture(s.pid).Wait()
}
