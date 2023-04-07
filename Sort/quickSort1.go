package main

import "fmt"

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pivot := nums[(left+right)/2]
	index := partition(nums, left, right, pivot)

	quickSort(nums, left, index-1)
	quickSort(nums, index, right)
}

func partition(nums []int, left, right int, pivot int) int {
	for left <= right {
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return left
}

func main() {
	// nums := []int{3, 2, 1, 5, 0, 2, 9, 8, 4}
	nums := []int{5, 6, 1, 8, -3, 0, 7, 95, 2}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
