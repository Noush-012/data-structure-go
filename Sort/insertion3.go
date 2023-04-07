package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	fmt.Println("Enter the elements, separated by spaces:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputSlice := strings.Split(input, " ")

	arr := make([]int, len(inputSlice))
	for i, v := range inputSlice {
		arr[i], _ = strconv.Atoi(v)
	}

	fmt.Println("Unsorted array is:", arr)
	insertionSort(arr)
	fmt.Println("Sorted array is: ", arr)
}
