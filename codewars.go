package main

import (
	"fmt"
	"math"
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
	//fmt.Println(WordsToMarks("attitude"))
	//fmt.Println(MakeNegative(0))
	//fmt.Println(NoSpace("   at tit ude"))
	//fmt.Println(SumMix([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7}))
	//fmt.Println(GetCount("abracadabra"))
	//fmt.Println(RowSumOddNumbers(13))
	//fmt.Println(GetSize(5,6,10))
	//fmt.Println(PartList([]string{"I", "wish", "I", "hadn't", "come"}))
	//fmt.Println(DNAStrand("ATTGC"))
	//fmt.Println(GetSum(5, -1))
	//fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 3))
	//fmt.Println(GetGrade(60, 82, 76))
	//fmt.Println(MxDifLg([]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}, []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}))
	//fmt.Println(TwiceAsOld(35,15))
	//fmt.Println(LoveFunc(5,9))
	//fmt.Println(Greet("Mark"))
	//fmt.Println(FindEvenIndex([]int{10, -80, 10, 10, 15, 35, 20}))
	//fmt.Println(Summation(50))
	//fmt.Println(findMostOftenRepeated([]int{110, 12, 35, 35, 20, 12, 35}))
	//fmt.Println(trimNegative(([]int{110, -12, 35, -35, 20, 12, -35})))
	//fmt.Println(trimLessAverage(([]int{53, -12, 70, -150, 99, 12, -35})))
	//    fmt.Println(insertArr([]int{1, 2, 3, 4, 5, 6, 7}, 8000, 2)) // добавление эл в произвольную позицию index
	//    fmt.Println(removeArr([]int{1, 2, 3, 4, 5, 6, 7}, 3)) // удаление эл в поизицию index
	//fmt.Println(MultiplicationTable(3))
	fmt.Println(SliseExpand([]int{1, 2, 3, 4, 5, 6, 7}, 5))
}

func SliseExpand(a []int, i int) []int {
	fmt.Println(a)
	a = append(a[:i], append(make([]int, 3), a[i:]...)...)
	return a
	// вставить 3 элемента в 5 позицию слайса
}

func MultiplicationTable(size int) [][]int {
	result := make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
		for j := 0; j < size; j++ {
			result[i][j] = (i + 1) * (j + 1)
		}
	}
	return result
	/* Создать таблицу умножения N×N размера, указанного в параметре.
			Например, если дано size3:
		1 2 3
		2 4 6
		3 6 9
		Для данного примера возвращаемое значение должно быть:
	[[1,2,3],[2,4,6],[3,6,9]]*/
}

func insertArr(list []int, el, index int) []int { // добавление эл в произвольную позицию index
	list = append(list, el)

	for key := len(list) - 1; key > index; key-- {
		list[key] = list[key-1]
		fmt.Println(list)
	}
	list[index] = el
	return list
}

func removeArr(list []int, index int) []int { // удаление эл в поизицию index
	for key := index + 1; key < len(list); key++ {
		list[key-1] = list[key]
		fmt.Println(list)
	}
	return list[:len(list)-1]
}

func trimLessAverage(array []int) []int {
	sum := 0
	for _, v := range array {
		sum += v
	}
	average := sum / len(array)
	newArr := make([]int, 0)
	for _, v := range array {
		if v > average {
			newArr = append(newArr, v)
		}
	}

	fmt.Println(newArr)
	return newArr
}

func trimNegative(array []int) []int {
	newArr := make([]int, 0)
	for _, v := range array {
		if v >= 0 {
			newArr = append(newArr, v)
		}
	}
	return newArr

}
func findMostOftenRepeated(array []int) (mostOften int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	counts := make(map[int]int)
	for i := 0; i < len(array); i++ {
		counts[array[i]]++
	}

	var maxCount = 0
	for key, value := range counts {
		if value > maxCount {
			maxCount = value
			mostOften = key
		}
	}
	return mostOften, nil

	/*
		var maxIndex, maxCount = 0, 0
		for i, number := range array {
			currentCount := 0
			for _, numberToCompare := range array {
				if number == numberToCompare {
					currentCount++
				}
			}

			if currentCount > maxCount {
				maxIndex = i
				maxCount = currentCount
			}
		}

		return array[maxIndex], nil*/
}

func Summation(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func FindEvenIndex(arr []int) int {
	sumLeft, sumRight := 0, 0
	for _, value := range arr {
		sumLeft = sumLeft + value
	}
	for index, value := range arr {
		sumLeft = sumLeft - value

		if sumLeft == sumRight {
			return index
		}
		sumRight = sumRight + value

	}
	return -1

	/*
			Вам будет дан массив целых чисел. Ваша задача состоит в том, чтобы взять этот массив и найти индекс N,
			где сумма целых чисел слева от N равна сумме целых чисел справа от N. Если такого индекса нет, верните -1.
		Например:
			Допустим, вам дан массив {1,2,3,4,3,2,1}:
			ваша функция вернет индекс 3, потому что в 3-й позиции массива сумма левой части индекса ( {1,2,3}) и сумма правой части индекса ( {3,2,1}) равны 6.
			Давайте посмотрим на другой.
				Вам дан массив {1,100,50,-51,1,1}:
			ваша функция вернет индекс 1, потому что в 1-й позиции массива сумма левой части индекса ( {1}) и сумма правой части индекса ( {50,-51,1,1}) равны 1.
		Последний:
			вам дан массив. {20,10,-80,10,10,15,35}
			В индексе 0 левая сторона равна {}
			Правая сторона равна {10,-80,10,10,15,35}
			Они обе равны 0при добавлении. (Пустые массивы равны 0 в этой задаче)
			Индекс 0 - это место, где левая и правая части равны.
				Примечание. Помните, что в большинстве языков программирования/скриптов индекс массива начинается с 0.
		Вход:
			Целочисленный массив длины 0 < arr < 1000. Числа в массиве могут быть любыми целыми положительными или отрицательными.
				Выход:
			наименьший индекс N, в котором сторона слева от Nравна стороне справа от N. Если вы не найдете индекс, соответствующий этим правилам, вы вернете -1.
			Примечание.
				Если вам дан массив с несколькими ответами, верните наименьший правильный индекс.*/
}

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s how are you doing today?", name)
}

func LoveFunc(flower1, flower2 int) bool {
	return (flower1+flower2)%2 != 0
}

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {

	return int(math.Abs(float64(sonYearsOld*2) - float64(dadYearsOld)))

	/*
		Your function takes two arguments:
		current father's age (years)
		current age of his son (years)
		Сalculate how many years ago the father was twice as old as his son (or in how many years he will be twice as old).
		The answer is always greater or equal to 0, no matter if it was in the past or it is in the future.
	*/

}

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	maxDiff := 0.0
	for _, x := range a1 {

		for _, y := range a2 {

			diff := math.Abs(float64(len(x) - len(y)))

			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}

	return int(maxDiff)
	/*
	   You are given two arrays a1 and a2 of strings.
	   Each string is composed with letters from a to z.
	   Let x be any string in the first array and y be any string in the second array.

	   		Find max(abs(length(x) − length(y)))

	   	If a1 and/or a2 are empty return -1 in each language except in Haskell (F#) where you will return Nothing (None).
	   		Example:
	   	a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
	   	a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
	   	mxdiflg(a1, a2) --> 13
	   	Bash note:
	   	input : 2 strings with substrings separated by ,
	   	output: number as a string
	*/
}

func GetGrade(a, b, c int) rune {
	score := (a + b + c) / 3
	switch {
	case 90 <= score && score <= 100:
		return 'A'
	case 80 <= score && score < 90:
		return 'B'
	case 70 <= score && score < 80:
		return 'C'
	case 60 <= score && score < 70:
		return 'D'
	default:
		return 'F'
	}

	//return rune("FFFFFFDCBAA"[(a+b+c)/30])
}

func LongestConsec(strarr []string, k int) string {
	n := len(strarr)
	var result string
	for i := 0; i <= n-k; i++ {
		temp := strings.Join(strarr[i:i+k], "")
		//fmt.Println(temp)
		if len(temp) > len(result) {
			result = temp
		}
	}
	return result
	/*You are given an array(list) strarr of strings and an integer k.
		Your task is to return the first longest string consisting of k consecutive strings taken in the array.
		strarr = ["tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"], k = 2

	Concatenate the consecutive strings of strarr by 2, we get:

	treefoling   (length 10)  concatenation of strarr[0] and strarr[1]
	folingtrashy ("      12)  concatenation of strarr[1] and strarr[2]
	trashyblue   ("      10)  concatenation of strarr[2] and strarr[3]
	blueabcdef   ("      10)  concatenation of strarr[3] and strarr[4]
	abcdefuvwxyz ("      12)  concatenation of strarr[4] and strarr[5]

	Two strings are the longest: "folingtrashy" and "abcdefuvwxyz".
	The first that came is "folingtrashy" so
	longest_consec(strarr, 2) should return "folingtrashy".

	In the same way:
	longest_consec(["zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"], 2) --> "abigailtheta"
	n being the length of the string array, if n = 0 or k > n or k <= 0 return "" (return Nothing in Elm, "nothing" in Erlang).

	Note
	consecutive strings : follow one after another without an interruption*/
}

func GetSum(a, b int) (res int) {
	if a == b {
		return a
	}
	max := math.Max(float64(a), float64(b))
	min := math.Min(float64(a), float64(b))
	for i := min; i <= max; i++ {
		res = res + int(i)
	}
	return res
}

func DNAStrand(dna string) string {
	res := ""

	for _, v := range dna {
		switch v {
		case 'A':
			res = res + "T"
		case 'T':
			res = res + "A"
		case 'G':
			res = res + "C"
		case 'C':
			res = res + "G"
		}
	}
	return res

}

func PartList(arr []string) string {
	//res := make([]string, 0)
	result := ""
	for i := 1; i < len(arr); i++ {
		firstPart := strings.Join(arr[:i], " ")
		secondPart := strings.Join(arr[i:], " ")
		//res = append(res, "("+firstPart+", "+secondPart+")")
		result = result + fmt.Sprintf("(%s, %s)", firstPart, secondPart)
	}
	//return strings.Join(res, "")
	return result
}

func GetSize(w, h, d int) [2]int {
	//sa := (2 * w * d) + (2 * h * d) + (2 * w * h)
	//v := w * h * d
	return [2]int{2 * (w*h + w*d + h*d), w * h * d}

}

func RowSumOddNumbers(n int) int {
	return n * n * n

	/*res := 0
	for i := 0; i < n; i++ {
		res = res + (n * n)
	}
	return res*/
	/*
		Given the triangle of consecutive odd numbers:
		           1
				3     5
		      7     9    11
		   13    15    17    19
		21    23    25    27    29
		...
		Calculate the sum of the numbers in the nth row of this triangle
		(starting at index 1) e.g.: (Input --> Output)
		1 -->  1
		2 --> 3 + 5 = 8
	*/
}

func GetCount(str string) (count int) {
	for _, v := range str {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	/*
		for _, v := range strings.Split(str, "") {
			if strings.Contains("aeiou", v) {
				count++
			}
		}
	*/
	return count

	//Return the number (count) of vowels in the given string.
	//We will consider a, e, i, o, u as vowels for this Kata (but not y).
	//The input string will only consist of lower case letters and/or spaces.

}

func SumMix(arr []interface{}) int {
	sum := 0

	for _, v := range arr {
		iv, _ := strconv.Atoi(fmt.Sprintf("%v", v))
		sum += iv

		/*
			switch v := v.(type) {
			case int:
				sum += v
			case string:
				num, _ := strconv.Atoi(v)
				sum += num
			} */

	}
	return sum
}

func NoSpace(word string) string {
	//return strings.ReplaceAll(word, " ", "")

	/*str := ""
	for _, v := range strings.Split(word, " ") {
		if !strings.Contains(v, " ") {
			str = str + v
		}
	}
	return str*/
	l := len(word)
	res, id := make([]rune, l), 0
	for j := 0; j < l; j++ {
		if word[j] != ' ' {
			res[id] = rune(word[j])
			id++
		}
	}

	return string(res[0:id])
}

func MakeNegative(x int) int {
	if x >= 0 {
		return -x
	}
	return x
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
