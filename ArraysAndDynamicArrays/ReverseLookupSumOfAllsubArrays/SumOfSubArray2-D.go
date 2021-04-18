package main

import "fmt"

// SumofSubArray finds sum of every possible subarray
// example SubArrays for [1,2,3] are [1], [2], [3], [1,2], [1,2,3], [2], [2,3], [3]
// Number of SubArrays for any array is N(N+1)/2
// Achieve in O(N)
// too many multiplication with higher values of arrays will lead to Integer Overflow
// which can be handled using taking modulo
// (a*b)%M = a%M * b%M wher M = 10^9+7
func SumofSubArray(arr [][]int) int {
	sum := 0
	// Numof rows
	N := len(arr)
	for i, row := range arr {
		for j, _ := range row {
			// No of elements in rows
			M := len(row)
			fmt.Println("i", sum)

			sum += (i + 1) * (j + 1) * arr[i][j] * (N - i) * (M - j)
		}

	}
	return sum
}

func main() {
	arr := [][]int{{1, 1}, {1, 1}}
	sum := SumofSubArray(arr)
	fmt.Println("Sum of SubArray", sum)
}
