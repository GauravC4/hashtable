package main

type BuiltinMap struct {
	cache map[int]int
}

func BuiltinMapConstructor() *BuiltinMap {
	bm := BuiltinMap{
		cache: make(map[int]int),
	}
	return &bm
}

func (bm *BuiltinMap) Set(key, value int) {
	bm.cache[key] = value
}

func (bm *BuiltinMap) Get(key int) (int, bool) {
	v, ok := bm.cache[key]
	if ok {
		return v, true
	}
	return 0, false
}

func (bm *BuiltinMap) Remove(key int) {
	delete(bm.cache, key)
}
