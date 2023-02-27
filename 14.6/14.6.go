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
	fmt.Println(result)
}
