package main

import "fmt"

func runningSum(nums []int) []int {
	var sum int = 0
	var result []int
	var index int = 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			sum += nums[i] + nums[j]
		}
		result[index] = sum
		index++
	}
	return result
}
func main() {
	var arr[] int
	fmt.Println("Enter Array")
	for i := 0; i < 5; i++ {
		fmt.Scan("%d", &arr[i])	
	}
	println(runningSum(arr))
}