package main

import (
	"fmt"
	"strings"
)

type PalindromeParam struct {
	sequence         string
	CASE_SENSITIVITY bool
}

func is_palindrome(options PalindromeParam) bool {
	sequence := options.sequence
	CASE_SENSITIVITY := options.CASE_SENSITIVITY
	if !CASE_SENSITIVITY {
		sequence = strings.ToLower(sequence)
	}
	reverse_string := []rune(sequence)
	for i, j := 0, len(sequence)-1; i < j; i, j = i+1, j-1 {
		reverse_string[i], reverse_string[j] = reverse_string[j], reverse_string[i]
	}
	if sequence == string(reverse_string) {
		return true
	}
	return false
}

func main() {
	var sequence string
	var CASE_SENSITIVITY bool
	var result bool
	for {
		fmt.Print("Enter a sequence: ")
		fmt.Scanln(&sequence)
		fmt.Print("Case Sensitivity [0/false | 1/true]: ")
		fmt.Scanln(&CASE_SENSITIVITY)
		result = is_palindrome(PalindromeParam{sequence: sequence, CASE_SENSITIVITY: CASE_SENSITIVITY})
		fmt.Println("Is", "\""+sequence+"\"", "palindrom with case-sensitivity", CASE_SENSITIVITY, ":", result, "\n-----------------")
	}
}
