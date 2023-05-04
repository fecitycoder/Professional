package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	a := [3]int{4, 5, 5}

	ref := a[:]
	fmt.Println("Existing array:\t", ref)
	t := append(s, ref...)
	fmt.Println("New slice:\t", t)
	s = append(s, ref...)
	fmt.Println("Existing slice:\t", s)
	s = append(s, s...)
	fmt.Println("s+s:\t\t", s)
}
