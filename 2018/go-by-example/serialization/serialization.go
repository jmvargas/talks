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
	group := Product{ID: 1, Name: "T-Shirt", Colors: []string{"Red", "Yellow", "Blue"}}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s", b)
}

// END OMIT
