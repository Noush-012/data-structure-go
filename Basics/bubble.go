package main

func bubbleSort(arr []int, data int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {

				arr[i], arr[j] = arr[j], arr[i]

			}

		}
	}

}

func main() {

}
