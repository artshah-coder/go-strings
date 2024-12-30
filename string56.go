// Given a string-sentence in Russian. Print the shortest word in the sentence.
// If there are several such words, then print the last one
package main

import (
	"errors"
	"fmt"
	"log"
	"unicode/utf8"
)

const STRING = "Строка-предложение на русском языке. Вывод самого короткого слова."

func main() {
	s, err := theShortestWord(STRING)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The shortest word in the string %s is %s\n", STRING, s)
}

func wordsFromLine(s string) ([]string, error) {
	strs := []string{}
	if len(s) == 0 {
		return strs, errors.New("Empty string!")
	}

	str := ""
	for _, ch := range s {
		switch {
		case ch >= 'А' && ch <= 'Я':
			str += string(ch)
		case ch >= 'а' && ch <= 'я':
			str += string(ch)
		case ch == 'ё' || ch == 'Ё':
			str += string(ch)
		default:
			if len(str) != 0 {
				strs = append(strs, str)
			}
			str = ""
		}
	}

	return strs, nil
}

func theShortestWord(s string) (string, error) {
	theShortest := s
	strs, err := wordsFromLine(s)
	if err != nil {
		return s, err
	}

	for _, str := range strs {
		if utf8.RuneCountInString(str) <= utf8.RuneCountInString(theShortest) {
			theShortest = str
		}
	}
	return theShortest, nil
}
