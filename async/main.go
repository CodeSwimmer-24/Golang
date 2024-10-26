package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	websiteList := []string{
		"https://www.youtube.com/",
		"https://www.google.com/",
		"https://digimazdoor.tech/",
	}

	var wg sync.WaitGroup

	for _, value := range websiteList {
		wg.Add(1)
		go getStatusCode(value, &wg)
	}

	wg.Wait()
}

func getStatusCode(endpoint string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("Error accessing %s: %v\n", endpoint, err)
	} else {
		fmt.Printf("Status code %d for %s\n", res.StatusCode, endpoint)
	}
}
