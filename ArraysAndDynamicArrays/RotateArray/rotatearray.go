package main

import "fmt"

// RotateArrayClockwiseBy1Unit rotates array by  1 unit
func RotateArrayClockwiseBy1Unit(arr []int) []int {

	N := len(arr)
	last := arr[N-1]
	for i := N - 2; i >= 0; i-- {
		arr[i+1] = arr[i]
	}
	arr[0] = last
	return arr
}

// RotateArrayClockwiseBykUnit rotates array by k unit
// Time complexity O(N)
// Space complexity O(N)
func RotateArrayClockwiseBykUnit(arr []int, k int) []int {

	N := len(arr)
	var result = make([]int, N)
	for i := 0; i <= N-1; i++ {
		result[(i+k)%N] = arr[i]
	}

	return result

}

func main() {
	arr := []int{1, 5, 2, 4, 3}
	// Rotate by 1 unit
	//result := RotateArrayClockwiseBy1Unit(arr)
	//fmt.Println("Rotated Array:", result)
	result2 := RotateArrayClockwiseBykUnit(arr, 3)
	fmt.Println("Rotated Array:", result2)
}
