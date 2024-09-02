package main

import (
	"bytes"
	"io"
	"testing"
)

func writeToBuffer(buf io.Writer, msg []byte) {
	buf.Write(msg)
}

func BenchmarkWriteToBuffer(b *testing.B) {
	msg := []byte("foo")
	buf := new(bytes.Buffer)

	for i := 0; i < b.N; i++ {
		for i := 0; i < 100; i++ {
			writeToBuffer(buf, msg)
      buf.Reset()
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	n := 100

	b.Run("slice non alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ints := []int{}
			for i := 0; i < n; i++ {
				ints = append(ints, i)
			}
		}
	})

	b.Run("slice alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ints := make([]int, n)
			for i := 0; i < n; i++ {
				ints[i] = i
			}
		}
	})
}
