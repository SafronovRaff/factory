package main

import "fmt"

/*Напишите функцию, которая определяет, является ли натуральное число палиндромом, и возвращает bool.*/

func main() {

	fmt.Println(palidrom(99099))
	//fmt.Println(palidrom(100))

}
func palidrom(num int) bool {
	a, b := 0, num
	for b > 0 {
		a = a * 10
		a = a + (b % 10)
		b = b / 10
	}
	return a == num
}
