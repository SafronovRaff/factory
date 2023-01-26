package main

import "fmt"

/*
Напишите функцию, которая ищет факториал натурального числа в итеративном и рекурсивном виде.

Факториал n — обозначается как n! — произведение чисел от 1 до n.

Например, факториал 4 — это 4! = 1 * 2 * 3 * 4 = 24
*/

func main() {

	fmt.Println(iner(4))
	fmt.Println(revers(4))
}

func iner(num int) int {
	sum := 1
	for i := 1; i <= num; i++ {
		sum = sum * i
	}
	return sum
}

func revers(num int) int {
	if num < 2 {
		return 1
	} else {
		return num * revers(num-1)
	}
}
