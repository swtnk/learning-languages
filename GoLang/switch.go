package main

import (
	"fmt"
	"time"
)

func main() {

	a := 1
	switch a {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Default")
	}

	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Its AM")
	case time.Now().Hour() == 12 && time.Now().Minute() == 00:
		fmt.Println("Its NOON")
	default:
		fmt.Println("Its PM")
	}

}
