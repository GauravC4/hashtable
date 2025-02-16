package main

type Hashable interface {
	Get(key int) (int, bool)
	Set(key, value int)
	Remove(key int)
}
