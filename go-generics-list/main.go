package main

import (
	"fmt"

	"github.com/zeindevs/go-generics-list/genericlist"
)

func main() {
	glist := genericlist.New[string]()

	glist.Insert("bob") // 0
	glist.Insert("foo") // 1
	glist.Insert("bar") // 2
	glist.Insert("baz") // 3

	fmt.Printf("%+v\n", glist)
	glist.Remove(1)
	fmt.Printf("%+v\n", glist)
	glist.RemoveByValue("baz")
	fmt.Printf("%+v\n", glist)
}
