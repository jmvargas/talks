package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

type User struct {
	ID   int    `json:"-"`
	Name string `json:"name,omitempty"`
	Age  int32  `json:"age"`
}

// START OMIT

func (u *User) canDrinkBeer() (bool, error) {
	if u.Age < 0 {
		return false, errors.New("Error: age lower than zero")
	}
	return u.Age >= 18, nil
}

func Handler(w http.ResponseWriter, r *http.Request) {
	user := User{ID: 1, Name: "Jesús", Age: 26}

	canDrinkBeer, err := user.canDrinkBeer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if canDrinkBeer {
		fmt.Fprintf(w, "%s, you can drink beer", user.Name)
	} else {
		fmt.Fprintf(w, "%s, you're so young!", user.Name)
	}
}

// END OMIT

func readBody(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}