package main

import (
	"fmt"
	"net/http"
)

var calls = 0

func main() {
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8080", nil)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	calls++
	fmt.Fprintf(w, "Hello World %d times! 👏😄", calls)
}
