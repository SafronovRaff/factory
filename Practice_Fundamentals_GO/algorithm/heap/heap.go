package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int // куча на основе среза

func (h IntHeap) Len() int {
	return len(h)
}

// Less - определяет логику выставления приоритета элементов
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Push -добавление нового элемента в кучу
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Swap - перестановка элдементов кучи
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Pop - извлечения элементов из кучи
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x

}
func main() {
	h := IntHeap{5, 12, 232}
	heap.Pop(h)

	fmt.Println(h)
}
