package main

import (
	"testing"
)

/*
inbuilt benchmark
set : 94.16 ns/op
get : 8.61 ns/op
remove : 54.66 ns/op
*/

/*
mymap benchmark
set : 137.6 ns/op
get : 44.87 ns/op
remove : 42.23 ns/op
*/

// var hashImplementation Hashable = BuiltinMapConstructor()
var hashImplementation Hashable = MyMapConstructor()

func TestSetGetRemove(t *testing.T) {
	arr := getRandArr(100)

	for _, v := range arr {
		hashImplementation.Set(v, v)
	}

	// test set and get
	for _, v := range arr {
		got, ok := hashImplementation.Get(v)
		if !ok || v != got {
			t.Errorf("did not find key %d", v)
		}
	}

	//fmt.Printf("TestSetGetRemove size after set %d\n", hashImplementation.Size())

	for _, v := range arr {
		hashImplementation.Remove(v)
	}

	//fmt.Printf("TestSetGetRemove size after remove %d\n", hashImplementation.Size())

	// test remove
	for _, v := range arr {
		_, ok := hashImplementation.Get(v)
		if ok {
			t.Errorf("found deleted key %d", v)
		}
	}

}

func TestSetGet(t *testing.T) {
	cases := []struct {
		key int
		val int
	}{
		{key: 23, val: 203},
		{key: 25, val: 205},
		{key: 23, val: 207},
		{key: 23, val: 209},
		{key: 27, val: 207},
	}

	for _, v := range cases {
		hashImplementation.Set(v.key, v.val)
		got, ok := hashImplementation.Get(v.key)
		if !ok || got != v.val {
			t.Errorf("did not get expected value %d for key %d", v.val, v.key)
		}
	}

}

func BenchmarkSet(b *testing.B) {
	n := b.N
	arr := getRandArr(n)

	b.ResetTimer()
	for i := 0; i < n; i++ {
		hashImplementation.Set(arr[i], arr[i])
	}
}

func BenchmarkGet(b *testing.B) {
	n := b.N
	arr := getRandArr(n * 2)
	for i := 0; i < n; i++ {
		hashImplementation.Set(arr[i], arr[i])
	}

	b.ResetTimer()
	for i := 0; i < n*2; i++ {
		hashImplementation.Get(arr[i])
	}
}

func BenchmarkRemove(b *testing.B) {
	n := b.N
	arr := getRandArr(n)
	for i := 0; i < n; i++ {
		hashImplementation.Set(arr[i], arr[i])
	}

	b.ResetTimer()
	for i := 0; i < n; i++ {
		hashImplementation.Remove(arr[i])
	}
}
