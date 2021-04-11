package main

import (
	"fmt"
)

func main() {

	condition := true

	if condition {
		fmt.Println("Its true")
	}

	if !condition {
		fmt.Println("Its not true")
	} else {
		fmt.Println("Its true again")
	}

	if name := "swetank"; name == "jonty" {
		fmt.Println("Name is jonty")
	} else if name == "swetank" {
		fmt.Println("Name is swetank")
	} else {
		fmt.Println("Name is not defined")
	}

}
