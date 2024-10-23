package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the file system")

	content := "Hello my name is fahad mahmood"

	file, error := os.Create("./myfile.txt")

	if error != nil {
		panic(error)
	} else {
		length, _ := io.WriteString(file, content)

		fmt.Println(length)
		defer file.Close()
	}
	readFile("./myfile.txt")
}

func readFile(filename string) {
	dataByte, error := ioutil.ReadFile(filename)
	if error != nil {
		panic(error)
	}
	fmt.Println(string(dataByte))
}
