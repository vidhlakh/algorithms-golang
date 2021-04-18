package main

import "fmt"

func getSum(i int, j int, arr []int) int {
	s := 0
	for k := i; k <= j; k++ {
		s += arr[k]

	}
	return s
}

// SumofSubArray finds sum of every possible subarray
// example SubArrays for [1,2,3] are [1], [2], [3], [1,2], [1,2,3], [2], [2,3], [3]
// Number of SubArrays for any array is N(N+1)/2
func SumofSubArray(arr []int) int {
	sum := 0
	N := len(arr)
	for i := 0; i <= N-1; i++ {
		for j := i; j <= N-1; j++ {
			sum += getSum(i, j, arr)
			fmt.Println("Sum", sum)
		}
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3}
	sum := SumofSubArray(arr)
	fmt.Println("Sum of SubArray", sum)
}
