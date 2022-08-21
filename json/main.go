package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	parseJsonFile()
}

func parseJsonFile() {
	// Open our jsonFile

	jsonFile, err := os.Open("charArr.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var arr []string
	_ = json.Unmarshal([]byte(byteValue), &arr)

	fmt.Println(arr[5])
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}

func indexOf(ch string, charArr []string) int {
	for i := 0; i < len(charArr); i++ {
		if ch == charArr[i] {
			return i
		}
	}

	return -1
}
