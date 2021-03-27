package main

import "fmt"

// PrefixSum finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func PrefixSum(a []int) []int {
	N := len(a)
	var prefixSum []int
	prefixSum = append(prefixSum, a[0])
	for i := 1; i < N; i++ {
		sum := prefixSum[i-1] + a[i]
		prefixSum = append(prefixSum, sum)
	}
	return prefixSum
}

// SuffixSum finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func SuffixSum(a []int) []int {
	N := len(a)
	var SuffixSum = make([]int, N)
	SuffixSum[N-1] = a[N-1]
	fmt.Println("Suff", SuffixSum)
	for i := N - 2; i >= 0; i-- {
		sum := SuffixSum[i+1] + a[i]
		SuffixSum[i] = sum
	}
	return SuffixSum
}
func main() {
	a := []int{2, 4, 10, 3, 7, 6, 3}
	result := PrefixSum(a)
	fmt.Println("Prefix Sum: ", result)
	result2 := SuffixSum(a)
	fmt.Println("Suffix Sum: ", result2)
	// use of Make make(type, length, capacity)
	//if we use append , it will append elements after length , so it is dynamic
	vals := make([]int, 5)
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
}
