package main

import (
	"fmt"
	"math"
)

// SumofSubArray finds sum of subarray
func SumofSubArray(arr []int, k int) {
	sum := 0
	N := len(arr)
	for i := 0; i <= k-1; i++ {
		sum += arr[i]
	}
	for j := k; j <= N-1; j++ {
		fmt.Println("Sum:", sum)
		sum += arr[j]   // add next elemnt of sliding window
		sum -= arr[j-k] // subtract previous element of sliding window
	}
	fmt.Println("Sum:", sum)
}

// FrequencyofElementofSubArray finds sum of subarray
// Frequency of eleement x in subarray
func FrequencyofElementofSubArray(arr []int, k int, x int) {
	freq := 0
	N := len(arr)
	for i := 0; i <= k-1; i++ {
		if arr[i] == x {
			freq++
		}
	}
	for j := k; j <= N-1; j++ {
		fmt.Println("Frequency:", freq)
		if arr[j] == x {
			freq++ // add next elemnt in fequency of sliding window
		}
		if arr[j-k] == x {
			freq-- // add next elemnt in fequency of sliding window
		}

	}
	fmt.Println("Frequencys:", freq)
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// CountsNumOfSwaps counts number of swaps required to bring all the elements <=k together
// arr[N] , K - find elements less than or equal to k grouped together
// First find number of elements less than or equal to k- windowSize
// Then Sliding window approach to find legal eleemnts (eleemmts less than k) on every window of size windowSize
func CountsNumOfSwaps(arr []int, k int) {
	windowSize := 0
	N := len(arr)
	// Find windowSize - number of elelmnts less that k
	for i := 0; i <= N-1; i++ {
		if arr[i] <= k {
			windowSize++
		}
	}
	maxCount := int(math.Inf(-1))
	legElements := 0

	for i := 0; i <= windowSize-1; i++ {
		if arr[i] <= k {
			legElements++
		}
		maxCount = Max(maxCount, legElements)
	}

	for j := windowSize; j <= N-1; j++ {
		if arr[j] <= k {
			legElements++
		}
		if arr[j-windowSize] <= k {
			legElements--
		}
		maxCount = Max(maxCount, legElements)

	}
	// swaps is WindowSize - MaxCount
	minSwaps := windowSize - maxCount
	fmt.Println("Minimum number of Swaps: ", minSwaps)
}
func main() {
	arr := []int{2, 7, 9, 5, 8, 7, 4}
	k := 5
	x := 2
	SumofSubArray(arr, k)
	FrequencyofElementofSubArray(arr, k, x)
	CountsNumOfSwaps(arr, k)
}
