package main

func fibonaciseries(n int) []int {
	var index = 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			sum := i + j
		}

		var result [index]int = sum
		index++

	}
	return result
}

// 0 1 1 2 3 5 7