package main

import "fmt"

// https://leetcode.com/problems/trapping-rain-water/
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the larger of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// PrefixMax finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func PrefixMax(a []int) []int {
	N := len(a)
	var prefixMax []int
	max := a[0]
	prefixMax = append(prefixMax, max)
	for i := 1; i < N; i++ {
		max = Max(prefixMax[i-1], a[i])
		prefixMax = append(prefixMax, max)
	}
	//fmt.Println("Pmax:", prefixMax)
	return prefixMax
}

// SuffixMax finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func SuffixMax(a []int) []int {
	N := len(a)
	var SuffixMax = make([]int, N)
	SuffixMax[N-1] = a[N-1]
	for i := N - 2; i >= 0; i-- {
		max := Max(SuffixMax[i+1], a[i])
		SuffixMax[i] = max
	}
	//fmt.Println("SuffixMax:", SuffixMax)
	return SuffixMax
}

// trap trapping rainwater problem
func trap(height []int) int {
	var result int
	N := len(height)
	if N == 0 {
		return result
	}
	pMax := PrefixMax(height)
	sMax := SuffixMax(height)
	for i := 1; i < N-1; i++ {
		//fmt.Println("hei", result, i, height[i])
		desiredHeight := Min(pMax[i-1], sMax[i+1])
		if desiredHeight > height[i] {
			result += desiredHeight - height[i]

		}
	}

	return result
}
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	output := trap(height)
	fmt.Println("Amount of water trapped:", output)
}
