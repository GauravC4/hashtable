package main

import "math/rand/v2"

func getRandArr(length int) []int {
	res := make([]int, length)
	r := rand.New(rand.NewPCG(4, 44))
	for i := range length {
		res[i] = r.IntN(2147483640)
	}
	return res
}
