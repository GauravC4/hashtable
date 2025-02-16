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

// change this to my implementation
var hashImplementation Hashable = BuiltinMapConstructor()

func TestSetGetRemove(t *testing.T) {
	arr := getRandArr(10)

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

	for _, v := range arr {
		hashImplementation.Remove(v)
	}

	// test remove
	for _, v := range arr {
		_, ok := hashImplementation.Get(v)
		if ok {
			t.Errorf("found deleted key %d", v)
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
