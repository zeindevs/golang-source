package consumer

import (
	"fmt"
	"reflect"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/zeindevs/microservices-hollywood/msg"
)

var processorPID = actor.NewPID("127.0.0.1:40000", "processor")

type Broker struct {
	accountID int64
}

func NewBroker(accId int64) actor.Producer {
	return func() actor.Receiver {
		return &Broker{
			accountID: accId,
		}
	}
}

// Receive implements actor.Receiver.
func (b *Broker) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Started:
		fmt.Println("[BROKER] started")
	default:
		fmt.Println(reflect.TypeOf(msg))
	}
}

func (b *Broker) OnStart(ctx *actor.Context) {
	go b.start(ctx)
}

func (b *Broker) OnStop(*actor.Context) {
	fmt.Println("[BROKER] stopped")
}

func (b *Broker) OnInit(*actor.Context) {
	fmt.Println("[BROKER] init")
}

func (b *Broker) start(ctx *actor.Context) {
	fmt.Println("[BROKER] started sending events to:", processorPID)
	for {
		time.Sleep(time.Second)
		ctx.Send(processorPID, &msg.PublicPayment{
			AccountID: b.accountID,
			Amount:    99.69,
		})
	}
}
