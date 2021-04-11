package main

import (
	"fmt"
)

func main() {

	a := make(map[string]string)
	a["first_name"] = "swetank"
	a["second_name"] = "subham"

	x, y := a["first_name"]
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(y)

	b := map[string]string{"first_name": "jonty", "second_name": "rhodes"}
	fmt.Println(b)

}
