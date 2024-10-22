package main

import "fmt"

func main() {
	// No inheratance is been allowed in GO

	details := User{"Fahad", "fahad@go.dev", true, 16}

	fmt.Println(details)
	fmt.Printf("Details %+v", details)
	fmt.Printf("Name %v \n", details.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
