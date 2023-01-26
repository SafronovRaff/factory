package main

import "fmt"

/*
Напишите программу, которая выводит на экран числа от 1 до 100. При этом вместо чисел, кратных трём,
программа должна выводить слово Fizz, а вместо чисел, кратных пяти — слово Buzz. Если число кратно и трём,
и пяти, то программа должна выводить слово FizzBuzz.
*/
func main() {

	for i := 1; i < 100; i++ {

		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%15 == 0 || i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}

	}

}
