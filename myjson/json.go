package myjson

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}

	return json.Marshal(s)
}

func TJson() {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

}

func JsonDecode() {

	//方式1
	const jsonStream = `
		[
			{"Name": "Ed", "Text": "Knock knock."},
			{"Name": "Sam", "Text": "Who's there?"},
			{"Name": "Ed", "Text": "Go fmt."},
			{"Name": "Sam", "Text": "Go fmt who?"},
			{"Name": "Ed", "Text": "Go fmt yourself!"}
		]
	`
	type Message struct {
		Name, Text string
	}
	var dst []map[string]string
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	if err := dec.Decode(&dst); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(dst)

	//方式2
	var dst1 []map[string]string
	json.Unmarshal([]byte(jsonStream), &dst1)
	fmt.Println(dst)
}
