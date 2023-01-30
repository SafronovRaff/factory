package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//fmt.Print(CountPositivesSumNegatives([]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -55, -13, -11, -12, -13, -14, -15}))
	//fmt.Print(MergeArrays([]int{1, 2, 3, 4}, []int{1, 6, 7, 8}))
	//mt.Println(century(201))
	//fmt.Println(Rps("rock", "scissors"))
	//fmt.Println(Derive(5, 9))
	//fmt.Println(Divisors(11))
	//fmt.Println(Invert([]int{1, 2, 3, 4, 5}))
	//fmt.Println(ReverseSeq(5))
	//fmt.Println(countSheep(2))
	//fmt.Println(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))
	//fmt.Println(CountBy(20, 5))
}

func CountBy(x, n int) []int {
	arr := make([]int, 0)
	for i := 1; i <= n; i++ {
		arr = append(arr, i*x)
	}
	return arr

	/*
		Create a function with two arguments that will return an array of the first n multiples of x.
		Assume both the given number and the number of times to count will be positive numbers greater than 0.
		Return the results as an array or list ( depending on language ).
		Examples
		countBy(1,10)  should return  {1,2,3,4,5,6,7,8,9,10}
		countBy(2,5)  should return {2,4,6,8,10}*/
}

func ToJadenCase(str string) string {

	//	return strings.Title(str) // устарела

	if len(str) == 0 {
		return ""
	}
	var sb strings.Builder
	split := strings.Split(str, " ")
	for _, v := range split {
		//strings.ToUpper(string(v[0]))
		fmt.Fprintf(&sb, strings.Replace(v, string(v[0]), strings.ToUpper(string(v[0])), 1)+" ")
	}
	return strings.Trim(sb.String(), "\n \r")

	/*words := strings.Split(str, " ")
	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		w := strings.ToUpper(string(word[0]))
		words[i] = w
		if len(word) == 1 {
			continue
		}
		words[i] = w + word[1:]
	}
	return strings.Join(words, " ")*/

}

func countSheep(num int) string {
	var sb strings.Builder
	for count := 1; count <= num; count++ {
		fmt.Fprintf(&sb, "%d shep...", count)
	}
	return sb.String()
	/*
		str := ""
		for i := 1; i <= num; i++ {
			num := strconv.Itoa(i)
			str += num + " sheep..."
		}
		return str
	*/
	/*Given a non-negative integer, 3 for example,
	return a string with a murmur: "1 sheep...2 sheep...3 sheep...".
	Input will always be valid, i.e. no negative integers.*/
}
func ReverseSeq(n int) []int {
	arr := []int{}
	for ; n != 0; n-- {
		arr = append(arr, n)
	}
	return arr
	/*
	   Build a function that returns an array of integers from n to 1 where n>0.
	   Example : n=5 --> [5,4,3,2,1]
	*/
}

func Invert(arr []int) []int {
	if len(arr) < 1 {
		return nil
	}
	invertArr := []int{}
	for _, v := range arr {
		v = v - (v * 2)
		invertArr = append(invertArr, v)
	}
	return invertArr

}

func Divisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count += 1
		}
	}

	return count
	/*	Count the number of divisors of a positive integer n.
		Random tests go up to n = 500000.
		Examples (input --> output)
		4 --> 3 (1, 2, 4)
		5 --> 2 (1, 5)
		12 --> 6 (1, 2, 3, 4, 6, 12)
		30 --> 8 (1, 2, 3, 5, 6, 10, 15, 30)*/
}

func Derive(coefficient, exponent int) string {
	//return fmt.Sprintf("%dx^%d ", coefficient*exponent, exponent-1)
	return strconv.Itoa(coefficient*exponent) + "x^" + strconv.Itoa(exponent-1)
}

func Rps(p1, p2 string) string {
	/*Rock Paper Scissors
	Let's play! You have to return which player won! In case of a draw return Draw!.
	Examples(Input1, Input2 --> Output):
	"scissors", "paper" --> "Player 1 won!"
	"scissors", "rock" --> "Player 2 won!"
	"paper", "paper" --> "Draw!"*/
	/*
		if p1 == "rock" && p2 == "rock" || p1 == "scissors" && p2 == "scissors" || p1 == "paper" && p2 == "paper" {
				return "Draw!"
			} else if p1 == "rock" && p2 == "scissors" || p1 == "scissors" && p2 == "paper" || p1 == "paper" && p2 == "rock" {
				return "Player 1 won!"
			} else {
				return "Player 2 won!"
			}
	*/
	var m = map[string]int{"rock": 0, "paper": 1, "scissors": 2}
	res := m[p1] - m[p2]
	if res == 0 {
		return "Draw!"
	} else if res == -1 || res == 2 {
		return "Player 2 won!"
	}
	return "Player 1 won!"

}

func century(year int) int {

	return (year + 99) / 100
	//return int(((year - 1) / 100)) + 1
	/*	res := 0
		for i := 0; i < year; i++ {
			if i%100 == 0 {
				res++
				fmt.Println(res)
			}
		}
		return res*/
	/* Introduction
	The first century spans from the year 1 up to and including the year 100,
	the second century - from the year 101 up to and including the year 200, etc.
		Task
	Given a year, return the century it is in.*/
}

func MergeArrays(arr1, arr2 []int) []int {
	arr3 := append(arr1, arr2...)
	keys := make(map[int]int)
	for _, v := range arr3 {
		keys[v] = v
		arr3 = append(arr3, v)
	}
	var arr4 []int
	fmt.Println(keys)
	for k := range keys {
		arr4 = append(arr4, k)
	}
	/* вариант через bool
	keys := make(map[int]bool)
		arr4 := []int{}
		for _, entry := range arr3 {
			if _, OK := keys[entry]; !OK {  // ок = ключ в мапе , если !ок
				keys[entry] = true 			// то добавим ключ значение
				arr4 = append(arr4, entry)
			}
	}
	*/
	sort.Ints(arr4)
	return arr4
}

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	if len(numbers) < 1 || numbers == nil {
		return res
	}
	res = []int{0, 0}
	for i, v := range numbers {
		if v > 0 {
			res[0] = i
		} else if v < 0 {
			res[1] += v
		}
	}
	return res
}
