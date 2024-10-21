package main

import "fmt"

func main() {

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["Ts"] = "TypeScripts"
	languages["Ry"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "Ry")
	fmt.Println(languages)

	// loops

	for key, values := range languages {
		fmt.Printf("For key %v, value %v\n", key, values)
	}

}
