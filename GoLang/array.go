package main

import (
	"fmt"
)

func main() {

	var a [3]int
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	fmt.Println(a)
	fmt.Println(a[1:3])

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(c)

}
