package main

import (
	"fmt"
	"sort"
)

func main() {
	// Initializing a slice of strings named 'friendsName' with three elements.
	var friendsName = []string{"Apple", "mango", "banana"}

	// Appending two more elements "Govava" and "Lichi" to the 'friendsName' slice.
	friendsName = append(friendsName, "Govava", "Lichi")

	// Printing the updated slice after appending new elements.
	fmt.Println(friendsName) // Output: [Apple mango banana Govava Lichi]

	// Slicing the 'friendsName' slice from index 1 to 3 (3 is exclusive).
	// This prints elements from index 1 to index 2: "mango", "banana".
	fmt.Println(friendsName[1:3]) // Output: [mango banana]

	// Creating an integer slice 'highScores' of length 4 using 'make'.
	highScores := make([]int, 4)

	// Assigning values to the elements of the slice.
	highScores[0] = 223
	highScores[1] = 263
	highScores[2] = 243
	highScores[3] = 213

	// Appending more values (322, 544, 655) to 'highScores' slice.
	highScores = append(highScores, 322, 544, 655)

	// Printing 'highScores' after appending new values.
	fmt.Println(highScores) // Output: [223 263 243 213 322 544 655]

	// Sorting the 'highScores' slice in ascending order.
	sort.Ints(highScores)
	// Printing the sorted 'highScores' slice.
	fmt.Println(highScores) // Output: [213 223 243 263 322 544 655]

	// Initializing a slice of strings named 'cources'.
	var cources = []string{"Java", "JS", "Ts", "Ruby", "SpringBoot"}

	// Printing the original 'cources' slice.
	fmt.Println(cources) // Output: [Java JS Ts Ruby SpringBoot]

	// Defining the index of the element to be removed (index 4, which is "SpringBoot").
	var index int = 4

	// Removing the element at index 4 by appending two parts of the slice:
	// 1. The slice before the element (cources[:index])
	// 2. The slice after the element (cources[index+1:])
	// This effectively skips the element at 'index'.
	cources = append(cources[:index], cources[index+1:]...)

	// Printing the 'cources' slice after removing the element at index 4.
	fmt.Println(cources) // Output: [Java JS Ts Ruby]
}
