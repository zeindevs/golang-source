package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
  Hello string `json:"hello"`
  ignored string
}

func main() {
	h := Message{Hello: "world"}
  b, _ := json.Marshal(h)
	fmt.Printf("%s\n", string(b))
}
