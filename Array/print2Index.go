package main

import "fmt"

func main() {
	i := 0
	j := 1
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	len := len(arr)
	fmt.Println(arr)
	for k := 0; k < len; k++ {
		if i < len && j < len {
			fmt.Println(arr[i], arr[j])
			i += 4
			j += 4
		} else if i == len-1 && j == len-1 {
			fmt.Println(arr[i])
			break
		} else {
			break
		}
	}

}
