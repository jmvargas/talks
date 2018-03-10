package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	re, _ := regexp.Compile("^/(.*)@privalia.com$")
	match := re.FindAllStringSubmatch(r.URL.Path, -1)

	if len(match) > 0 {
		fmt.Fprintf(w, "Hello %s! 👋", match[0][1])
		return
	}
	fmt.Fprintf(w, "Sorry, I don't know you")
}
