package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func withID(id string) OptFunc {
	return func(o *Opts) {
		o.id = id
	}
}

func WithTLS(opt *Opts) {
	opt.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(o *Opts) {
		o.maxConn = n
	}
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	s := newServer(WithTLS, withMaxConn(90), withID("foo"))
	fmt.Printf("%+v\n", s)
}
