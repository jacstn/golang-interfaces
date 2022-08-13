package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	people := `
	[{
		"first_name": "john",
		"last_name": "travolta"
	}]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(people), &unmarshalled)
	if err != nil {
		log.Println("error handling json")
	}

	fmt.Printf("unmarshalled: %v", unmarshalled)

	marshalled, err := json.MarshalIndent(unmarshalled, "", "   ")
	if err != nil {
		log.Println("error marshalling json")
	}
	print(string(marshalled))
}
