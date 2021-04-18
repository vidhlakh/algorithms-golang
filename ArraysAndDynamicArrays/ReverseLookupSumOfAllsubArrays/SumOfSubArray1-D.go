package main

import "fmt"

// SumofSubArray finds sum of every possible subarray
// example SubArrays for [1,2,3] are [1], [2], [3], [1,2], [1,2,3], [2], [2,3], [3]
// Number of SubArrays for any array is N(N+1)/2
// Achieve in O(N)
// too many multiplication with higher values of arrays will lead to Integer Overflow
// which can be handled using taking modulo
// (a*b)%M = a%M * b%M wher M = 10^9+7
func SumofSubArray(arr []int) int {
	sum := 0
	N := len(arr)
	for i := 0; i <= N-1; i++ {
		sum += (i + 1) * arr[i] * (N - i)
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3}
	sum := SumofSubArray(arr)
	fmt.Println("Sum of SubArray", sum)
}
