package bar

import "fmt"

type BarInterface interface {
	SayHello()
}

type Bar struct{}

func (b *Bar) SayHello() {
	fmt.Println("Hello!")
}

func (b *Bar) Sayhola() {
	fmt.Println("Hola!")
}

func (b *Bar) SayPrivet() {
	fmt.Println("Privet!")
}

func (b *Bar) SayGoddag() {
	fmt.Println("God dag!")
}
