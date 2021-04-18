package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//IsArmStrong displays whether a number is Armstrong number or not
func IsArmStrong(input int) (output int) {
	r := 0

	for input > 0 {
		r = input % 10
		output += r * r * r
		input = input / 10
	}
	return output
}
func main() {
	var testCases int
	var input, output int
	var inputArr []string
	scanner := bufio.NewScanner(os.Stdin)

	t := scanner.Scan()
	if t {
		testCases, _ = strconv.Atoi(scanner.Text())
	}
	for i := 0; i < testCases; i++ {
		scanner.Scan()
		inputArr = append(inputArr, scanner.Text())
	}
	for _, data := range inputArr {
		input, _ = strconv.Atoi(data)
		output = IsArmStrong(input)
		if output == input {
			fmt.Println(input, "is an Armstrong Number")
		} else {
			fmt.Println(input, "is not an Armstrong Number")
		}

	}

}
