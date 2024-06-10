package consumer

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
	"github.com/zeindevs/microservices-hollywood/msg"
)

type PublicProcessor struct {
}

func NewPublicProcessor() actor.Producer {
	return func() actor.Receiver {
		return &PublicProcessor{}
	}
}

func (p *PublicProcessor) Receive(ctx *actor.Context) {
	switch m := ctx.Message().(type) {
	case *msg.PublicPayment:
		fmt.Printf("[PUBLIC PROC] %v\n", m)
	}
}
