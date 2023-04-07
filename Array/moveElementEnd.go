package main

import "fmt"

func moveElementToEnd(array []int, target int) []int{
    i := 0
    j := len(array) - 1
    for i < j {
        for i < j && array[i] != target {
            i++
        }
        for i < j && array[j] == target {
            j--
        }
        if i < j {
            array[i], array[j] = array[j], array[i]
        }
    }
    return array
}

func main() {
    arr := []int{1, 3, 4, 4, 5, 9, 6, 6}
    result  := moveElementToEnd(arr, 4)
    fmt.Println(result)
}

