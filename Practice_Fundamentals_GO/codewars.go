package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Print(CountPositivesSumNegatives([]int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11, -55, -13, -11, -12, -13, -14, -15}))
	//fmt.Print(MergeArrays([]int{1, 2, 3, 4}, []int{1, 6, 7, 8}))
	//mt.Println(century(201))
	//fmt.Println(Rps("rock", "scissors"))

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
