package main

import "fmt"

func main() {
	// No inheratance is been allowed in GO

	details := User{"Fahad", "fahad@go.dev", true, 16}

	fmt.Println(details)
	fmt.Printf("Details %+v", details)
	fmt.Printf("Name %v \n", details.Name)

	details.getStatus()
	details.setMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println(u.Status)
}

func (e User) setMail() {
	e.Email = "newFahad@gmail.com"
	fmt.Println(e.Email)
}
