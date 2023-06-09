package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{
			Telephone{Mobile: true, Number: "123-456"},
			Telephone{Mobile: true, Number: "123-abcd"},
			Telephone{Mobile: false, Number: "abcd-456"},
		},
	}
	saveToJSON(os.Stdout, myRecord)
}
