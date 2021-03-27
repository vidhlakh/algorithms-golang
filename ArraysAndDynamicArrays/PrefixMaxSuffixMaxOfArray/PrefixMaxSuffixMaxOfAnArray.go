package main

import (
	"fmt"
)

// Max returns the larger of x or y.
func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int64) int64 {
	if x > y {
		return y
	}
	return x
}

// PrefixMax finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func PrefixMax(a []int64) []int64 {
	N := len(a)
	var prefixMax []int64
	max := a[0]
	prefixMax = append(prefixMax, max)
	for i := 1; i < N; i++ {
		max = Max(prefixMax[i-1], a[i])
		prefixMax = append(prefixMax, max)
	}
	return prefixMax
}

// SuffixMax finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func SuffixMax(a []int64) []int64 {
	N := len(a)
	var SuffixMax = make([]int64, N)
	SuffixMax[N-1] = a[N-1]
	for i := N - 2; i >= 0; i-- {
		max := Max(SuffixMax[i+1], a[i])
		SuffixMax[i] = max
	}
	return SuffixMax
}
func main() {
	a := []int64{2, 4, 10, 3, 7, 6, 3}
	result := PrefixMax(a)
	fmt.Println("Prefix Max: ", result)
	result2 := SuffixMax(a)
	fmt.Println("Suffix Max: ", result2)

}
