package main

import "fmt"

func reverse(arr []int) {
    start := 0
    end := len(arr) - 1
    for start < end {
        arr[start], arr[end] = arr[end], arr[start]
        start++
        end--
    }
}

func main() {
    var n int
    fmt.Print("Enter the size of array: ")
    fmt.Scanf("%d", &n)
    arr := make([]int, n)
    for i := 0; i <len(arr); i++ {
		fmt.Printf("Enter the %d th element of array ", i+1)
		fmt.Scanf("%d", &arr[i])
	} 
    reverse(arr)
    fmt.Println(arr) 
}
