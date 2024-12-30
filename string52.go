// Given a string sentence in Russian. Convert the string so that each word begins with
// a capital letter. Words that do not begin with a letter must not be changed
package main

import (
	"errors"
	"fmt"
	"log"
)

const STRING = ",,,строка-предложение на русском языке. капитализация всех слов строки..."

func main() {
	s, err := CapLine(STRING)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Source line: %s.\nCapitalized line: %s\n", STRING, s)
}

func CapLine(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("Empty string!")
	}

	str := ""
	isNW := true
	for _, ch := range s {
		if isNW {
			switch {
			case ch >= 'а' && ch <= 'я':
				isNW = false
				str += string(ch - 32)
				continue
			case ch == 'ё':
				isNW = false
				str += "Ё"
				continue
			case ch == ' ':
				str += string(ch)
				continue
			default:
				isNW = false
			}
		} else if ch == ' ' {
			isNW = true
		}
		str += string(ch)
	}

	return str, nil
}
