package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := make([]int, 2)
	for i := 0; i < 2; i++ {
		var value int
		fmt.Print("Enter Value: ")
		fmt.Scanln(&value)
		b = append(b, value)
	}

	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b[2:4]))

}
