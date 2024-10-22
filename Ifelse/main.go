package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := 10

	var result string

	if num > 10 {
		result = "The number is greater than 10"
	} else if num < 10 {
		result = "The number is less than 10"
	} else {
		result = "The number is equal to 10"
	}
	fmt.Println(result)

	diceNumber := rand.Intn(6) + 1

	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		{
			fmt.Println("Your value is 1")
		}
	case 2:
		{
			fmt.Println("Your value is 2")
		}
	case 3:
		{
			fmt.Println("Your value is 3")
		}
	case 4:
		{
			fmt.Println("Your value is 4")
		}
	case 5:
		{
			fmt.Println("Your value is 5")
		}
	case 6:
		{
			fmt.Println("Your value is 6")
		}
	}
}
