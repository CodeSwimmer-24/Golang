package main

import "fmt"

func main() {
	result := adder(2, 5)
	fmt.Println(result)

	resultAll := addAll(1, 2, 4, 5)
	fmt.Println(resultAll)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func addAll(val ...int) int {
	total := 0
	for _, num := range val {
		total += num
	}
	return total
}
