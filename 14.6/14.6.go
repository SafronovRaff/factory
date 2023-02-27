package main

import "fmt"

func main() {
	// исходная длинна слайсов
	lenArrOne, lemArrTwo := 0, 0
	fmt.Println("введите размер первого массива")
	fmt.Scanln(&lenArrOne)
	fmt.Println("введите размер второго массива")
	fmt.Scanln(&lemArrTwo)
	fmt.Println("введите значения первого массива")

	// инициализируем первый слайс
	arrOne := make([]string, 0)

	// цикл заполнения первого слайса
	for i := 0; i < lenArrOne; i++ {
		s := ""
		fmt.Scanln(&s)
		arrOne = append(arrOne, s)
	}
	fmt.Println("введите значения второго массива")
	// инициализируем второй слайс
	arrTwo := make([]string, 0)

	// цикл заполнения второго слайса
	for i := 0; i < lemArrTwo; i++ {
		s := ""
		fmt.Scanln(&s)
		arrTwo = append(arrTwo, s)
	}
	// инициализируем мапу ключ строки, значения пустая структура
	dictionary := make(map[string]struct{})

	//итерация по первому слайсу, для каждого элемента добавляем запись в
	//мапу, где ключ сам элемент, а значение пустая структура

	for _, v := range arrOne {
		dictionary[v] = struct{}{}
	}
	// инициализируем строковый слайс
	var result []string

	// итерация по второму слайсу с проверкой значения на наличие этого
	// значения в мапе
	for _, v := range arrTwo {
		// если элемент присутствует, то добавляем его в result
		if _, ok := dictionary[v]; ok {
			result = append(result, v)
		}
	}

	// тестовое
	v := []int{3, 3, 4, 5, 8, 9, 10, 10, 15, 15}
	repit(v)

}

// создаём пустую мапу и слайс для результата
// итерация по входящему массиву с добавлением значения ключом мапы и значение мапы +1
// итерация ключа и значения мапы
// если значение = 1, то записываем ключ в слайс с результатом
func repit(arr []int) {
	dictionary := make(map[int]int)
	result := make([]int, 0)
	for _, v := range arr {
		dictionary[v]++
	}
	for key, value := range dictionary {
		if value == 1 {
			result = append(result, key)
		}
	}
	fmt.Print(result)
}
