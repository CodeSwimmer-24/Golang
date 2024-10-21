package main

import "fmt"

func main() {
    // Print an introductory message
    fmt.Println("Study about pointer")

    // Declare a pointer variable, but it is not initialized (nil pointer)
    var pointer *int

    // Print the value of the uninitialized pointer (it will print <nil>)
    fmt.Println("Value of pointer:", pointer)

    // Declare a variable `num` and assign a value to it
    num := 23

    // Create a pointer `ptr` that stores the memory address of `num`
    var ptr = &num

    // Print the pointer (which holds the memory address of `num`)
    fmt.Println("Memory address stored in ptr:", ptr)

    // Dereference the pointer to get the value stored at the memory address (which is 23)
    fmt.Println("Value pointed to by ptr:", *ptr)

	*ptr = *ptr + 2;
	fmt.Println("New value:", num)
}
