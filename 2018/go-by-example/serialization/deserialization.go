package main

import (
	"encoding/json"
	"fmt"
)

// START OMIT

type Product struct {
	ID     int      `json:"-"`
	Name   string   `json:"name,omitempty"`
	Colors []string `json:"colors"`
}

func main() {
	myJSON := `{"name":"T-Shirt","colors":["Red","Yellow","Blue"]}`
	product := &Product{}
	err := json.Unmarshal([]byte(myJSON), product)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%v", product)
}

// END OMIT
