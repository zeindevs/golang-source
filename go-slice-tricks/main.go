package main

import (
	"fmt"
	"sort"
)

type Numbers []int

type byInc struct {
	Numbers
}

func (n byInc) Len() int           { return len(n.Numbers) }
func (n byInc) Swap(i, j int)      { n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i] }
func (n byInc) Less(i, j int) bool { return n.Numbers[i] < n.Numbers[j] }

type byDec struct {
	Numbers
}

func (n byDec) Len() int           { return len(n.Numbers) }
func (n byDec) Swap(i, j int)      { n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i] }
func (n byDec) Less(i, j int) bool { return n.Numbers[i] > n.Numbers[j] }

func (n Numbers) Len() int           { return len(n) }
func (n Numbers) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Numbers) Less(i, j int) bool { return n[i] < n[j] }

func main() {
	numbers := Numbers{1, 10, 4, 9, 3}
	fmt.Println(numbers)

	// sort.Sort(numbers)
	sort.Sort(byInc{numbers})

	fmt.Println(numbers)

	sort.Sort(byDec{numbers})

	fmt.Println(numbers)
}

type MySlice []int

func (s MySlice) Remove(index int) []int {
	return append(s[:index], s[index+1:]...)
}

func sample2() {
	numbers := MySlice{1, 2, 3, 4, 5}
	numbers = numbers.Remove(1)

	fmt.Println(removeFromSlice(numbers, 1))
	fmt.Println(removeFromSliceWithOrder(numbers, 1))
}

func removeFromSliceWithOrder(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func removeFromSlice(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

var users = []string{}

func addUser(user ...string) {
	users = append(users, user...)
}

func sample() {
	addUser("alice", "bob", "foo")
	fmt.Println(users)
}
