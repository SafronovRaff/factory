package main

import "fmt"

/*Ещё одна задачка с собеседования. Напишите функцию — простейший архиватор.
	Она должна принимать на вход непустую строку и преобразовывать её так,
чтобы повторяющиеся идущие подряд буквы указывались один раз с количеством повторений.

Например: “aaaa” => “a4” или “abcccccc” => “a1b1c6”.
*/

func main() {
	fmt.Println(arhiv("aaa***GGeeelll"))
}

func arhiv(str string) string {
	var letter rune
	letterLen := 0
	var result string

	for _, value := range str {
		if value != letter {
			if letterLen > 0 {
				result = result + fmt.Sprintf("%c%d", letter, letterLen)
			}
			letter = value
			letterLen = 1
		} else {
			letterLen = letterLen + 1
		}
	}
	result = result + fmt.Sprintf("%c%d", letter, letterLen)
	return result
}
