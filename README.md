Custom implementation of hashmap in golang.
For simplicity used int key value pairs only and used chaining method for collision resolution.

inbuilt benchmark
set : 94.16 ns/op
get : 8.61 ns/op
remove : 54.66 ns/op

mymap benchmark
set : 137.6 ns/op
get : 44.87 ns/op
remove : 42.23 ns/op

set and remove are close, get is 5x slow ! maybe open addressing is faster for get due to spatial locality ... 