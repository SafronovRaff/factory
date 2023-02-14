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
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	ar = mergeSort(ar)
	fmt.Println(ar)

}

func mergeSort(ar []int) []int {
	n := len(ar)
	if n < 2 {
		return ar
	}

	mid := n / 2
	left := mergeSort(ar[:mid])
	right := mergeSort(ar[mid:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	res := make([]int, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
