package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/anthdm/hollywood/actor"
)

// 1. make it work
// 2. make it better
// 3. make it faster

type StationData struct {
	name string
	temp float64
}

type Station struct {
	min, max, avg float64
}

type Compute struct {
	stations map[string]*Station
}

func NewCompute() actor.Receiver {
	return &Compute{
		stations: map[string]*Station{},
	}
}

func (r *Compute) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case StationData:
		_ = msg
	}
}

type Aggregator struct {
	computePID *actor.PID
	workers    int
}

func NewAggregator(cpid *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &Aggregator{
			computePID: cpid,
		}
	}
}

func (a *Aggregator) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case []byte:
		ctx.SpawnChildFunc(func(c *actor.Context) {
			switch c.Message().(type) {
			case actor.Stopped:
				a.workers--
			case actor.Started:
				a.workers++
				r := bufio.NewReader(bytes.NewReader(msg))
				for {
					b, err := r.ReadBytes('\n')
					if err != nil {
						if errors.Is(err, io.EOF) {
							break
						}
						log.Fatal(err)
					}
					line := strings.TrimRight(string(b), "\n")
					parts := strings.Split(line, ";")
					station := parts[0]
					temp, err := strconv.ParseFloat(parts[1], 64)
					if err != nil {
						log.Fatal(err)
					}
					ctx.Send(a.computePID, StationData{
						name: station,
						temp: temp,
					})
				}
				// process
				c.Engine().Poison(c.PID())
			}
		}, "foert")
	}
}

func main() {
	defer func(start time.Time) {
		fmt.Printf("took us %v\n", time.Since(start))
	}(time.Now())

	// e, err := actor.NewEngine(actor.NewEngineConfig())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// computePid := e.Spawn(NewCompute, "compute")
	// pid := e.Spawn(NewAggregator(computePid), "agg")

	f, err := os.Open("data/measurements.txt")
	if err != nil {
		log.Fatal(err)
	}

	chunkDataCh := make(chan []byte, 2048)

	chunkSize := 128 * 1024 * 1024
	buf := make([]byte, chunkSize)
	leftover := make([]byte, 0, chunkSize)

	go func() {
		for {
			n, err := f.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				panic(err)
			}
			buf = buf[:n]
			finalBuf := make([]byte, n)
			copy(finalBuf, buf)

			lastNewLineIndex := bytes.LastIndex(buf, []byte{'\n'})

			finalBuf = append(leftover, buf[:lastNewLineIndex+1]...)
			leftover = make([]byte, len(buf[lastNewLineIndex+1:]))
			copy(leftover, buf[lastNewLineIndex+1:])

			chunkDataCh <- finalBuf
		}
	}()

	for chunk := range chunkDataCh {
		go processChunk(chunk)
	}
}

func processChunk(chunk []byte) {
	fmt.Println(chunk)
}
