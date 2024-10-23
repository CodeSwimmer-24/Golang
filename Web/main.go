package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const urlView = "https://jsonplaceholder.typicode.com/posts"

func main() {

	response, err := http.Get(urlView)

	panicMode(err)

	fmt.Printf("Type of response %T \n", response)

	data, err := ioutil.ReadAll(response.Body)

	panicMode(err)

	readData := string(data)

	fmt.Print(readData)

	checkUrl()

	defer response.Body.Close()

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "jsonplaceholder.typicode.com",
		Path:   "/posts",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}

func checkUrl() {
	result, err := url.Parse(urlView)
	panicMode(err)

	fmt.Println(result.Host)
}

func panicMode(err error) {
	if err != nil {
		panic(err)
	}
}
