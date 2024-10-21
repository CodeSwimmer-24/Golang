package main

import "fmt"

// Define a constant for a login token
const LoginToken = "sdfghjkl678cgyuiokm"

func main() {
	// Print "hello world" to the console
	fmt.Println("hello world")

	// Declare and initialize a string variable 'name'
	var name string = "Fahad"
	// Print the type of 'name' variable
	fmt.Printf("My name is: %T \n", name)

	// Declare and initialize a 64-bit integer variable 'num'
	var num int64 = 9886789654
	// Print the value of 'num' using println (prints a newline automatically)
	println(num)
	// Print the type of 'num' variable
	fmt.Printf("My name is: %T \n", num)

	// Declare an integer variable 'newNum' and assign a value to it
	var newNum int
	newNum = 344
	// Print the value of 'newNum'
	println(newNum)

	// Perform an addition operation on 'newNum' and store in 'addVal'
	var addVal = newNum + newNum
	// Print the result of the addition
	fmt.Println(addVal)

	// Declare and initialize a string variable using shorthand syntax
	myname := "fahad"
	// Print the value of 'myname'
	fmt.Println(myname)

	// Print the constant 'LoginToken'
	fmt.Println(LoginToken)
}
