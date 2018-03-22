package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// START OMIT

type User struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) makeYounger() {
	u.Age--
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	user := User{}
	json.Unmarshal(readBody(r), &user)
	user.makeYounger()
	data, _ := json.Marshal(user)
	fmt.Fprintf(w, "%s", data)
}

// END OMIT

func readBody(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}
