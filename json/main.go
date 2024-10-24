package main

import (
	"encoding/json"
	"fmt"
)

type cources struct {
	Name     string `json:"courceName"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// JsonBulinding()
	DecodeJson()
}

func JsonBulinding() {

	allCources := []cources{
		{
			"React Js", 399, "freecodecamp", "abc123", []string{"Web Dev", "Js"},
		},
		{
			"Angular", 499, "freecodecamp", "zsed123", []string{"Web Dev", "Js"},
		},
		{
			"React Native", 599, "freecodecamp", "wedf123", nil,
		},
	}

	finalJson, err := json.MarshalIndent(allCources, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
	{
                "courceName": "React Js",
                "Price": 399,
                "Platform": "freecodecamp",
                "tags": [
                        "Web Dev",
                        "Js"
                ]
        }
	`)

	var courcesName cources

	checkValidJson := json.Valid(jsonData)

	if checkValidJson {
		fmt.Println("data is valid")
		json.Unmarshal(jsonData, &courcesName)
		fmt.Printf("%#v\n", courcesName)
	} else {
		fmt.Println("Not data is valid")
	}

	var keyValuePair map[string]interface{}
	json.Unmarshal(jsonData, &keyValuePair)
	fmt.Printf("%#v\n", keyValuePair)

	for k, v := range keyValuePair {
		fmt.Printf("The key is %v and the value is %v \n", k, v)
	}
}
