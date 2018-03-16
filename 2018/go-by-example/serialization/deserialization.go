package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// START OMIT

type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.Unmarshal(readBody(r), &user)
	fmt.Fprintf(w, "Hello %s, your are %d years old! 👋", user.Name, user.Age)
}

// END OMIT

func readBody(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}
