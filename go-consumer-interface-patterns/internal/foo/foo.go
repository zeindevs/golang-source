package foo

type bar interface {
	SayHello()
}

type Foo struct {
	bar bar
}

func NewFoo(bar bar) *Foo {
	return &Foo{
		bar: bar,
	}
}

func (f *Foo) Greet() {
	f.bar.SayHello()
}
