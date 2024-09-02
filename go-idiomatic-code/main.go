package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

// constant declarations
const (
	Scalar  = 0.1
	Version = 0.1
)

// variables grouping
func Foo() int {
	var (
		x   = 100
		y   = 2
		foo = "foo"
	)
	fmt.Println(foo)

	return x + y
}

// functions that panic
func MustParseIntFromString(s string) int {
	// Logic
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

// struct initialization
type Vector struct {
	x int
	y int
}

// mutext grouping
type Server struct {
	listenAddr string
	isRunning  bool

	othersMu sync.RWMutex
	others   map[string]net.Conn

	peersMu sync.RWMutex
	peers   map[string]net.Conn
}

// interface declarations/naming
type Getter interface {
	Get()
}

type Putter interface {
	Put()
}

type Storer interface {
	Getter
	Putter
}

// function grouping
func VeryImportantFuncExported() {}

func veryImportantFunc() {}

func simpleUtil() {}

// http handler naming
func handleGetUserById() {}

func handleResizeImage() {}

// enums (kinda!!)
type Suit byte

const (
	SuitHarts Suit = iota
	SuitClubs
	SuitDiamonds
	SuitSpades
)

// constructor
type Order struct {
	Size float64
}

// func NewOrder(size float64) *Order
func New(size float64) *Order {
	return &Order{
		Size: size,
	}
}

func main() {
	_ = Vector{
		x: 10,
		y: 20,
	}
}
