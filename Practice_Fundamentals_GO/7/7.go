package main

import "fmt"

/*Напишите программу, выводящую первые 20 простых чисел.
Простое число — натуральное число больше 1, которое делится без остатка только на себя и на 1.
*/

func main() {
	fmt.Println(2)

	count := 1
	for i := 3; count < 20; i++ {
		c := 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				c++
			}
		}
		if c == 0 {
			fmt.Println(i)
			count = count + 1
		}
	}
}
