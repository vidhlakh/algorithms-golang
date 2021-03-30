package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/max-chunks-to-make-sorted/
// Input: arr = [4,3,2,1,0]
// Output: 1
// Explanation:
// Splitting into two or more chunks will not return the required result.
// For example, splitting into [4, 3], [2, 1, 0] will result in [3, 4, 0, 1, 2], which isn't sorted.

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

//maxChunksToSorted Max chunks
func maxChunksToSorted(arr []int) int {
	var result int
	N := len(arr)
	MAX := int(math.Inf(-1))
	for i := 0; i <= N-1; i++ {
		MAX = Max(arr[i], MAX)
		if MAX == i {
			result++
		}
	}
	return result
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5}
	output := maxChunksToSorted(arr)
	fmt.Println("Maximum Chunks:", output)
}
