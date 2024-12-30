// Given a string - binary representation of a positive integer. Print a string, representing
// the decimal notation of the same number.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"unicode/utf8"
)

func main() {
	f := os.Stdin
	defer f.Close()

	fmt.Print("Input a binary notation of integer positive number (for example: 10101): ")
	var source string
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		source = scanner.Text()
	} else {
		fmt.Println("\nWhile reading your input an error occured!\nBye!")
		return
	}

	if s, err := convertToDecimal(source); err != nil {
		fmt.Println("While handling your input the error occured:", err)
		return
	} else {
		fmt.Println("Inputed number in the decimal notation: ", s)
	}
}

func convertToDecimal(str string) (string, error) {

	intRes := 0
	var i int
	for i = 0; i < utf8.RuneCountInString(str); i++ {
		switch str[i] {
		case '1':
			intRes += int(math.Pow(2, float64(utf8.RuneCountInString(str)-i-1)))
		case '0':
			continue
		default:
			return "", errors.New("inputed string is not " +
				"a binary notation of integer positive number!")
		}
	}

	var bs []byte
	for i = 0; intRes != 0; i++ {
		bs = append(bs, byte(intRes%10)+'0')
		intRes /= 10
	}

	for i = 0; i < utf8.RuneCount(bs)/2; i++ {
		bs[i], bs[utf8.RuneCount(bs)-i-1] = bs[utf8.RuneCount(bs)-i-1], bs[i]
	}

	return string(bs), nil
}
