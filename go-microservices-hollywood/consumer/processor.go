package consumer

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
	"github.com/zeindevs/microservices-hollywood/msg"
)

type Processor struct {
	publicProcPID *actor.PID
}

func NewProcessor() actor.Producer {
	return func() actor.Receiver {
		return &Processor{}
	}
}

func (p *Processor) Receive(ctx *actor.Context) {
	switch m := ctx.Message().(type) {
	case *msg.PublicPayment:
		ctx.Forward(p.publicProcPID)
	default:
		fmt.Println(m)
	}
}

func (p *Processor) OnStart(ctx *actor.Context) {
	p.publicProcPID = ctx.Engine().Spawn(NewPublicProcessor(), "public")
	fmt.Println(p.publicProcPID)
}

func (p *Processor) OnStop(ctx *actor.Context) {
	ctx.Engine().Poison(p.publicProcPID)
}

func (p *Processor) OnInit(*actor.Context) {
	panic("unimplemented")
}
