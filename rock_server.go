package main

import (
	"db/people"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", rockers)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func rockers(response http.ResponseWriter, request *http.Request) {
	people := people.FindAll()
	fmt.Fprint(response, people)
}
