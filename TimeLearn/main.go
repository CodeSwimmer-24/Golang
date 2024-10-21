package main

import (
	"fmt"
	"time"
)

func main() {
	// Welcome message
	fmt.Println("Welcome to my timezone")

	// Get the current time
	todayTime := time.Now()

	// Print the current time
	fmt.Println(todayTime)

	// Format the current time using the proper layout (01-02-2006 15:04:05 is the reference layout)
	formattedTime := todayTime.Format("01-02-2006 15:04:05 Monday")

	// Print the formatted time
	fmt.Println(formattedTime)
}
