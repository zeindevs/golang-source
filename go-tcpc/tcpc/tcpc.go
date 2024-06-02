package tcpc

import (
	"encoding/gob"
	"log"
	"net"
	"sync"
	"time"
)

// func NewChannel[T any]) (*Sender[T], *Receiver[T])

type TCPC[T any] struct {
	listenAddr string
	remoteAddr string

	Sendchan     chan T
	Recvchan     chan T
	outboundConn net.Conn
	ln           net.Listener
	wg           sync.WaitGroup
}

func New[T any](listenAddr, remoteAddr string) (*TCPC[T], error) {
	tcpc := &TCPC[T]{
		listenAddr: listenAddr,
		remoteAddr: remoteAddr,
		Sendchan:   make(chan T, 10),
		Recvchan:   make(chan T, 10),
	}

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return nil, err
	}
	tcpc.ln = ln

	go tcpc.loop()
	go tcpc.acceptLoop()
	go tcpc.dialRemoteAndRead()

	return tcpc, nil
}

func (t *TCPC[T]) loop() {
	t.wg.Wait()

	for {
		msg := <-t.Sendchan
		if err := gob.NewDecoder(t.outboundConn).Decode(&msg); err != nil {
			log.Println(err)
			return
		}
		log.Printf("sending msg over the wire to %s: %v\n", t.remoteAddr, msg)
	}
}

func (t *TCPC[T]) acceptLoop() {
	defer func() {
		t.ln.Close()
	}()

	for {
		conn, err := t.ln.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			return
		}

		log.Printf("sender connected %s", conn.RemoteAddr())

		go t.handleConn(conn)
	}
}

func (t *TCPC[T]) handleConn(conn net.Conn) {
	defer func() {
		conn.Close()
	}()

	for {
		var msg T
		if err := gob.NewDecoder(conn).Decode(&msg); err != nil {
			log.Println(err)
			return
		}
		t.Recvchan <- msg
	}
}

func (t *TCPC[T]) dialRemoteAndRead() {
	t.wg.Add(1)

	for {
		conn, err := net.Dial("tcp", t.remoteAddr)
		if err != nil {
			log.Printf("dial error (%s)\n", err)
			time.Sleep(3 * time.Second)
		} else {
			t.outboundConn = conn
			break
		}
	}

	t.wg.Done()
}
