package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// PrefixMax finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func PrefixMax(a []int, p int, N int) []int {

	var prefixMax []int
	max := p * a[0]
	prefixMax = append(prefixMax, max)
	for i := 1; i < N; i++ {
		max = Max(prefixMax[i-1], p*a[i])
		prefixMax = append(prefixMax, max)
	}
	return prefixMax
}

// SuffixMax finds prefix sum of an array a[0-1], a[0-2], a[0-3],,,a[0-n]
func SuffixMax(a []int, r int, N int) []int {

	var SuffixMax = make([]int, N)
	SuffixMax[N-1] = r * a[N-1]
	for i := N - 2; i >= 0; i-- {
		max := Max(SuffixMax[i+1], r*a[i])
		SuffixMax[i] = max
	}
	return SuffixMax
}

// Solve finds the whole expression
func Solve(a []int, pMax []int, sMax []int, q int, N int) int {

	result := int(math.Inf(-1))
	for i := 0; i < N; i++ {
		x := pMax[i] + q*a[i] + sMax[i]
		result = Max(x, result)
	}
	return result
}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

// main expression maximize the expression p*a[i]+q*a[j]+r*a[k]
func main() {
	// p := 1
	// q := 2
	// r := 3
	//var a = []int{1, 2, 3, 4, 5}
	// Give stdin n,p,q ,r (5 1 2 3 )
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	data := numbers(scanner.Text())

	n, p, q, r := data[0], data[1], data[2], data[3]
	// Give stdin array (1 2 3 4 5)
	scanner.Scan()
	a := numbers(scanner.Text())

	pMax := PrefixMax(a, p, n)
	sMax := SuffixMax(a, r, n)
	result := Solve(a, pMax, sMax, q, n)
	fmt.Println(result)
}
