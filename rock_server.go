package main

import (
	"db/people"
	"fmt"
)

func main() {
	people := people.FindAll()
	fmt.Println(people)
}
