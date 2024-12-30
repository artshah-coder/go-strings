// Given a string consisting of Russian words, typed in capital letters and separated
// by spaces (one or more). Find the number of words that begin and end with the same letter
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Анаконда Ананас Шашлык Тибет Воронов"
	strs := strings.Fields(s)
	count := 0
	for _, str := range strs {
		str = strings.ToLower(str)
		var first, last rune
		for _, ch := range str {
			first = ch
			break
		}
		if last, _ = utf8.DecodeLastRuneInString(str); last == first {
			count++
		}
	}
	fmt.Println("Result:", count)
}
