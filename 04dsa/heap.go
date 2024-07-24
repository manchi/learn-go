package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println("heap test")

	nums := &intHeap{3, 1, 4, 5, 1, 1, 2, 6}
	heap.Init(nums)

	fmt.Println(nums)
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}
