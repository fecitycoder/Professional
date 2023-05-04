package main

import "runtime"

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var sctructure []data
	for i := 0; i < N; i++ {
		value := int(i)
		sctructure = append(sctructure, data{value, value})
	}

	runtime.GC()
	_ = sctructure[0]
}
