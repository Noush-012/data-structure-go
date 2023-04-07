package main

import "fmt"

func twoNumberSum(arr []int, target int) []int {
for i := 0; i < len(arr)-1; i++ {
for j := i + 1; j < len(arr); j++ {
if arr[i]+arr[j] == target {
return []int{arr[i], arr[j]}
}
}
}
return []int{}
}
func main() {
array := []int{6, 5, 4, 7, 9, 0, 2}
target := 10
result := twoNumberSum(array, target)
for _, v := range result {
fmt.Println(v)
}
}