package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	myRecord := Record{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []Telephone{
			Telephone{Mobile: true, Number: "123-456"},
			Telephone{Mobile: true, Number: "123-abcd"},
			Telephone{Mobile: false, Number: "abcd-123"},
		}}

	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rec))

	var unRec Record
	err1 := json.Unmarshal(rec, &unRec)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(unRec)
}
