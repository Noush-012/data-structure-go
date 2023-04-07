package main

import "fmt"

func binarySearch(arr []int, key int) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == key {
            return mid
        } else if arr[mid] < key {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1
}

func main() {
    arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    key := 6
    index := binarySearch(arr, key)
    if index != -1 {
        fmt.Println("Element found at index:", index)
    } else {
        fmt.Println("Element not found")
    }
}
