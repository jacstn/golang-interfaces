package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	arr := parseJsonFile()

	fmt.Println(indexOf("A", arr))
	en := IntToCode(23423423423, arr)
	fmt.Println("encode int", en)
	fmt.Println("dec", CodeToInt(en, arr))
}

func IntToCode(num int, charArr []string) string {
	arrLen := len(charArr)
	fmt.Println("arrlen", arrLen)
	var strOut string

	for num > arrLen-1 {
		idx := (num % int(arrLen))
		fmt.Println("current index", idx)
		strOut = charArr[idx] + strOut
		num = int(math.Floor(float64(num / arrLen)))
	}

	return charArr[num%arrLen] + strOut
}

func CodeToInt(code string, charArr []string) int {
	retVal := 0

	for i := 0; i < len(code); i++ {
		x := float64(len(charArr))
		y := float64(len(code) - i - 1)
		retVal += int(math.Pow(x, y)) * indexOf(string(code[i]), charArr)
	}
	return retVal
}

func parseJsonFile() []string {
	// Open our jsonFile

	jsonFile, err := os.Open("charArr.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var arr []string
	_ = json.Unmarshal([]byte(byteValue), &arr)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	return arr

}

func indexOf(ch string, charArr []string) int {
	for i := 0; i < len(charArr); i++ {
		if ch == charArr[i] {
			return i
		}
	}

	return -1
}
