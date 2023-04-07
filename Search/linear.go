package main

import "fmt"

func linearSearch(arr []int, key int) int {
    for i, v := range arr {
        if v == key {
            return i
        }
    }
    return -1
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    key := 9
    index := linearSearch(arr, key)
    if index != -1 {
        fmt.Println("Element found at index:", index)
    } else {
        fmt.Println("Element not found")
    }
}
