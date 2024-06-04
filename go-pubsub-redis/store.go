package main

import (
	"fmt"
)

type Store struct {
	data  map[int]string
	cache Cacher
}

func NewStore(c Cacher) *Store {
	data := map[int]string{
		1: "Elon Musk is the new owner of Twitter",
		2: "Foo is not bar and bar is not bar",
		3: "Must watch AnthonyGG",
	}
	return &Store{
		data:  data,
		cache: c,
	}
}

func (s *Store) Get(key int) (string, error) {
	val, ok := s.cache.Get(key)
	if ok {
		// busting the cache
		// if err := s.cache.Remove(key); err != nil {
		// 	fmt.Print(err)
		// }
		fmt.Println("returning the value from cache")
		return val, nil
	}

	val, ok = s.data[key]
	if !ok {
		return "", fmt.Errorf("key not found: %d", key)
	}

	if err := s.cache.Set(key, val); err != nil {
		return "", err
	}

	fmt.Println("returning key from internal storage")

	return val, nil
}
