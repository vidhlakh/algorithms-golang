package main

import (
	"fmt"
	"math"
)

// PrefixMax finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func PrefixMax(a []int) []int {
	N := len(a)
	var prefixMax []int
	prefixMax = append(prefixMax, math.Inf(-1))
	for i := 1; i < N; i++ {
		sum := prefixMax[i-1] + a[i]
		prefixMax = append(prefixMax, sum)
	}
	return prefixMax
}

// SuffixMax finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func SuffixMax(a []int) []int {
	N := len(a)
	var SuffixMax = make([]int, N)
	SuffixMax[N-1] = a[N-1]
	fmt.Println("Suff", SuffixMax)
	for i := N - 2; i >= 0; i-- {
		sum := SuffixMax[i+1] + a[i]
		SuffixMax[i] = sum
	}
	return SuffixMax
}
func main() {
	a := []int{2, 4, 10, 3, 7, 6, 3}
	result := PrefixMax(a)
	fmt.Println("Prefix Max: ", result)
	result2 := SuffixMax(a)
	fmt.Println("Suffix Max: ", result2)
	// use of Make make(type, length, capacity)
	//if we use append , it will append elements after length , so it is dynamic
	vals := make([]int, 5)
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
}
