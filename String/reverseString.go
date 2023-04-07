package main

import "fmt"

func reverseString(s string) string {
    var reversedString string
    for i := len(s) - 1; i >= 0; i-- {
        reversedString += string(s[i])
    }
    return reversedString
}

func main() {
    originalString := "Welcome You"
    reversedString := reverseString(originalString)
    fmt.Println("Original String:", originalString)
    fmt.Println("Reversed String:", reversedString)
}
