package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	//SortingByChoice(ar)
	bidirectionalSelectionSort(ar)
	fmt.Println(ar)
}

func SortingByChoice(ar []int) {
	n := len(ar)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if ar[j] > ar[minIndex] {
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}
func bidirectionalSelectionSort(ar []int) {
	n := len(ar)
	left := 0
	right := n - 1

	for left < right {
		minIndex := left
		maxIndex := right
		for i := left; i <= right; i++ {

			if ar[i] < ar[minIndex] {
				minIndex = i
			}
			if ar[i] > ar[maxIndex] {
				maxIndex = i
			}
		}
		ar[left], ar[minIndex] = ar[minIndex], ar[left]
		if left == maxIndex {
			maxIndex = minIndex
		}
		ar[right], ar[maxIndex] = ar[maxIndex], ar[right]
		left++
		right--
	}

}
