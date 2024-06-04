package main

type Cacher interface {
	Get(key int) (string, bool)
	Set(key int, val string) error
	Remove(key int) error
}

type NopCache struct{}

func (c *NopCache) Get(key int) (string, bool) {
	return "", false
}

func (c *NopCache) Set(key int, val string) error {
	return nil
}

func (c *NopCache) Remove(int) error {
	return nil
}
