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

	insertionSort(ar)

	fmt.Println(ar)
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
