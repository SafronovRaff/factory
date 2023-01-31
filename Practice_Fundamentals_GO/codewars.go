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
	//fmt.Println(Move(3, 5))
	//fmt.Println(TwoToOne("bsjвавамсbjs", "sdттндддsds"))
	//fmt.Println(StringToArray("str string uhuh ssds"))
	//fmt.Println(FakeBin("338392736587"))
	//fmt.Println(Between(4, 6))
	//fmt.Println(IsTriangle(8, 3, 5))
	//fmt.Println(GetMiddle("123456"))
	//fmt.Println(FindMultiples(5, 25))
	fmt.Println(WordsToMarks("attitude"))
}

func WordsToMarks(s string) int {
	res := 0
	/*for _, value := range s {
		if value >= 'a' && value <= 'z' {
			res = res + int(value-'a') + 1
		}
	}	*/

	for _, value := range s {
		res = res + int(value) - 96 //ASCII: a = 97,..., z = 122
	}
	return res

	/*
		a = 1, b = 2, c = 3 ... z = 26
		Then l + o + v + e = 54
		and f + r + i + e + n + d + s + h + i + p = 108
		So friendship is twice as strong as love :-)
		Your task is to write a function which calculates the value of a word based off the sum of the alphabet positions of its characters.
		The input will always be made of only lowercase letters and will never be empty.*/
}
func FindMultiples(integer, limit int) []int {
	/*var res []int
	num := integer
	for i := integer; i <= limit; i = i + num {
		res = append(res, i)
	}*/

	res := make([]int, limit/integer+1)
	for i, j := integer, 0; i <= limit; i, j = i+integer, j+1 {
		res[j] = i
	}

	return res[:limit/integer]
}
func GetMiddle(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	return s[(n-1)/2 : n/2+1]

	/*if len(s)%2 == 0 {
		sl := s[(len(s)/2)-1 : len(s)/2+1]
		return sl
	}
	return s[(len(s) / 2) : (len(s)/2)+1]*/

	/*You are going to be given a word. Your job is to return the middle character of the word.
	If the word's length is odd, return the middle character.
	If the word's length is even, return the middle 2 characters.
	#Examples:
	Kata.getMiddle("test") should return "es"
	Kata.getMiddle("testing") should return "t"
	Kata.getMiddle("middle") should return "dd"
	Kata.getMiddle("A") should return "A"*/
}

func IsTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && c+b > a

	/*Implement a function that accepts 3 integer values a, b, c.
	The function should return true if a triangle can be built with the sides of given length and false in any other case.
	(In this case, all triangles must have surface greater than 0 to be accepted).*/
}

func Between(a, b int) []int {
	// Time: 2737 ms
	/*res := make([]int, 0)
	for a <= b {
		res = append(res, a)
		a++
	}
	return res*/

	//Time: 1503ms
	res := make([]int, b-a+1)
	for i := a; i <= b; i++ {
		res[i-a] = i
	}
	return res
}

func FakeBin(x string) string {
	/*	res := ""
		for _, v := range x {
			num := int(v - '0')
			if num < 5 {
				res = res + "0"
			} else {
				res += "1"
			}
		}
		return res*/
	/*	Строка num := int(v - '0') преобразует строковый символ v в целое число. Вот как это работает: v — это символ в строке x,
		поэтому его тип — rune (что в Go эквивалентно char в других языках программирования).
		«0» — это символьный литерал, представляющий код ASCII для цифры 0. Вычитая «0» из v, мы получаем разницу между их кодами ASCII.
		Эта разница будет представлять собой целочисленное значение цифры, представленной v. Наконец, функция int преобразует полученную разность (типа rune) в тип int.*/

	bs := []byte(x)
	for i, v := range bs {
		if v < '5' {
			bs[i] = '0'
		} else {
			bs[i] = '1'
		}
	}
	return string(bs)

	/*Given a string of digits, you should replace any digit below 5 with
	'0' and any digit 5 and above with '1'. Return the resulting string.*/
}

func StringToArray(str string) []string {
	//return strings.Split(str, " ")
	return strings.Fields(str)
}

func TwoToOne(s1 string, s2 string) string {
	/*str := strings.Split(s1+s2, "")
	sort.Strings(str)
	res := ""
	for _, v := range str {
		chr := string(v)
		if !strings.Contains(res, chr) {
			res = res + chr
		}
	}
	return res*/
	str := s1 + s2
	result := ""
	for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
		if strings.Contains(str, char) {
			result += char
		}
	}
	return result

	/*
		str := s1 + s2
		buf := make(map[string]bool)
		res := []string{}
		for _, v := range str {
			if unicode.Is(unicode.Latin, v) {
				buf[string(v)] = true
			}
		}
		for key, _ := range buf {
			res = append(res, key)
		}
		sort.Strings(res)
		return strings.Join(res, "")*/

	/*
		Take 2 strings s1 and s2 including only letters from a to z.
		Return a new sorted string, the longest possible,
		containing distinct letters - each taken only once - coming from s1 or s2.
			  Examples:
			  a = "xyaabbbccccdefww"
			  b = "xxxxyyyyabklmopq"
			  longest(a, b) -> "abcdefklmopqwxy"

			  a = "abcdefghijklmnopqrstuvwxyz"
			  longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"*/
}

func Move(position int, roll int) int {
	return position + (roll * 2)
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
