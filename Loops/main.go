package main

import "fmt"

func main() {

	days := []string{"Sunday", "Monday", "Tudeday", "wednesday", "Thursday", "Friday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, days := range days {
		fmt.Printf("index is %v and value is %v \n", index, days)
	}

}
