package main

import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	testString := "MALAYALAM"
	result := isPalindrome(testString)
	if result {
		fmt.Println(testString, "is a palindrome.")
	} else {
		fmt.Println(testString, "is not a palindrome.")
	}
}
