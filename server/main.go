package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	getApiData()
	postApiData()
	forDataPost()
}

func getApiData() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// Make the GET request
	response, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("Error making the GET request:", err)
		return
	}
	// Ensure the response body is closed after the function finishes
	defer response.Body.Close()

	// Check if the response status is OK
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Received non-OK status code:", response.StatusCode)
		return
	}

	// Read the response body
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	var responseString strings.Builder

	finalStat, _ := responseString.Write(data)

	fmt.Println(finalStat, responseString.String())
	// Print the response data
	// fmt.Println(string(data))
}

func postApiData() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	// fake json Payload

	requestBody := strings.NewReader(`
	{
		"country":"india",
		"courency":"inr"
	}
	`)
	response, _ := http.Post(myUrl, "application/json", requestBody)

	data, _ := io.ReadAll(response.Body)

	fmt.Print(string(data))

}

func forDataPost() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts"

	data := url.Values{}

	data.Add("firstName", "Fahad")
	data.Add("lastName", "mahmood")
	data.Add("email", "dev@shc.co.in")

	response, _ := http.PostForm(myUrl, data)

	defer response.Body.Close()

	realData, _ := io.ReadAll(response.Body)

	fmt.Print(string(realData))
}
