package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Welcome message
	welcome := "Welcome to our shop"
	fmt.Println(welcome)

	// Create a new reader to capture user input
	reader := bufio.NewReader(os.Stdin)

	// Ask for a rating input
	fmt.Println("Please give your rating:")

	// Read input from user until a newline character is encountered
	inputValue, _ := reader.ReadString('\n')

	// Trim any extra spaces/newline from input
	inputValue = strings.TrimSpace(inputValue)

	// Thank user for the input
	fmt.Println("Thanks for rating us:", inputValue)

	// Convert the string input to a float
	numRating, err := strconv.ParseFloat(inputValue, 64)

	// Error handling in case conversion fails
	if err != nil {
		fmt.Println("There was an error processing your rating. Please enter a valid number.")
	} else {
		// Add 1 to the rating as per the original logic
		fmt.Println("Your adjusted rating is:", numRating+1)
	}
}
