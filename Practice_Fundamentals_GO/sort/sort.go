package sort

import (
	"math/rand"
)

func bubbleSort(ar []int) {
	l := len(ar)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
}

func selectionSort(ar []int) {
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

func insertionSort(ar []int) {
	n := len(ar)
	for i := 1; i < n; i++ { //Внешний цикл начинается со второго элемента массива и продолжается до конца.
		key := ar[i] // Переменный ключ устанавливается на текущий рассматриваемый элемент
		j := i - 1
		for j >= 0 && ar[j] > key { //Внутренний цикл выполняет итерацию по отсортированному на данный момент подмассиву, сдвигая элементы вправо до тех пор,
			ar[j+1] = ar[j] //пока не будет найдено правильное положение для ключа.
			j--
		}
		ar[j+1] = key
	}
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

func quickSort(ar []int) {
	if len(ar) < 2 {
		return
	}
	left, right := 0, len(ar)-1
	povorIndex := rand.Int() % len(ar)
	ar[povorIndex], ar[right] = ar[right], ar[povorIndex]
	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}
	ar[left], ar[right] = ar[right], ar[left]
	quickSort(ar[:left])
	quickSort(ar[left+1:])
	return
}
