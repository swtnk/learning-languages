package main

import (
	"fmt"
	"reflect"
)

func main() {

	// CONTANTS
	const x int = 5
	const y = 5
	const z string = "Hello"
	fmt.Println("---------CONSTANTS---------")
	fmt.Println(reflect.TypeOf(x), x)
	fmt.Println(reflect.TypeOf(y), y)
	fmt.Println(reflect.TypeOf(z), z)

	var a int
	a = 4
	var k int8
	k = 127
	var l uintptr
	l = 12512512451215412118
	m := 32768
	d := 5
	var n int

	// INTEGERS
	fmt.Println("---------INTEGERS---------")
	fmt.Println(reflect.TypeOf(a), a)
	fmt.Println(reflect.TypeOf(k), k, k+1, k+2)
	fmt.Println(reflect.TypeOf(l), l)
	fmt.Println(reflect.TypeOf(m), m)
	fmt.Println(reflect.TypeOf(d), d)
	fmt.Println(reflect.TypeOf(n), n)

	var b string
	b = "Swetank"
	e := "string"
	var o string

	// STRINGS
	fmt.Println("---------STRINGS---------")
	fmt.Println(reflect.TypeOf(b), b)
	fmt.Println(reflect.TypeOf(e), e)
	fmt.Println(reflect.TypeOf(o), o)

	var c rune
	c = 'S'
	var p rune

	// RUNES
	fmt.Println("---------RUNES---------")
	fmt.Println(reflect.TypeOf(c), c)
	fmt.Println(reflect.TypeOf(c), c, string(c))
	fmt.Println(reflect.TypeOf(p), p, string(p))
	fmt.Println(string(c) == "S")

	var i float32
	i = 4.8
	var j float64
	j = 9.2
	f := 5.2

	// FLOATS
	fmt.Println("---------FLOATS---------")
	fmt.Println(reflect.TypeOf(i), i)
	fmt.Println(reflect.TypeOf(j), j)
	fmt.Println(reflect.TypeOf(f), f)

	var g bool
	g = true
	h := false

	// BOOLEANS
	fmt.Println("---------BOOLEANS---------")
	fmt.Println(reflect.TypeOf(g), g)
	fmt.Println(reflect.TypeOf(h), h)

}
