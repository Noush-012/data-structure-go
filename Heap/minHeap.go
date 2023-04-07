package main

import "fmt"

type MinHeap struct {
	data []int
}

func (h *MinHeap) ParentIndex(i int) int {
	return (i - 1) / 2
}

func (h *MinHeap) LeftChildIndex(i int) int {
	return 2*i + 1
}

func (h *MinHeap) RightChildIndex(i int) int {
	return 2*i + 2
}

func (h *MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) UpHeapify(i int) {
	for h.data[i] < h.data[h.ParentIndex(i)] {
		h.Swap(i, h.ParentIndex(i))
		i = h.ParentIndex(i)
	}
}

func (h *MinHeap) DownHeapify(i int) {
	for {
		leftChildIndex := h.LeftChildIndex(i)
		rightChildIndex := h.RightChildIndex(i)

		// If both child indices are out of bounds, we're done
		if leftChildIndex >= len(h.data) && rightChildIndex >= len(h.data) {
			break
		}

		// Find the minimum child index
		minChildIndex := leftChildIndex
		if rightChildIndex < len(h.data) && h.data[rightChildIndex] < h.data[leftChildIndex] {
			minChildIndex = rightChildIndex
		}

		// If the parent is already smaller than both children, we're done
		if h.data[i] <= h.data[minChildIndex] {
			break
		}

		// Otherwise, swap with the smaller child and continue down the heap
		h.Swap(i, minChildIndex)
		i = minChildIndex
	}
}

func (h *MinHeap) Insert(x int) {
	h.data = append(h.data, x)
	h.UpHeapify(len(h.data) - 1)
}

func (h *MinHeap) ExtractMin() (int, error) {
	if len(h.data) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	// Extract the minimum element and replace it with the last element
	min := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	// Fix the heap property by down-heapifying
	h.DownHeapify(0)

	return min, nil
}

func main() {
	h := &MinHeap{}
	h.Insert(3)
	h.Insert(2)
	h.Insert(1)
	h.Insert(4)
	h.Insert(5)

	for len(h.data) > 0 {
		min, err := h.ExtractMin()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(min)
	}
}
