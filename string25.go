// Given a string representing the decimal notation of a positive integer.
// Print the binary representation of the same number.
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
	var f *os.File
	f = os.Stdin
	defer f.Close()

	fmt.Print("Input a decimal integer positive number: ")
	var source string
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		source = scanner.Text()
	} else {
		fmt.Println("\nWhile reading your input an error occured!\nBye!")
	}

	if s, err := convertToBinary(source); err != nil {
		fmt.Println("While handling your input the error occured: ", err)
		return
	} else {
		fmt.Println("Inputed number as the binary: ", s)
	}
}

func convertToBinary(str string) (string, error) {

	intRes := 0

	for i := 0; i < utf8.RuneCountInString(str); i++ {
		switch {
		case str[i] >= '0' && str[i] <= '9':
			intRes += int((str[i] - '0')) * int(math.Pow10(utf8.RuneCountInString(str)-i-1))
		default:
			return "", errors.New("inputed string is not a decimal integer positive number")
		}
	}

	var bs []byte
	for i := 0; intRes > 0; i++ {
		if i == 0 {
			bs = make([]byte, 1)
			bs[i] = byte(intRes%2) + '0'
		} else {
			bs = append(bs, byte(intRes%2)+'0')
		}
		intRes /= 2
	}

	for i := 0; i < utf8.RuneCount(bs)/2; i++ {
		bs[i], bs[utf8.RuneCount(bs)-i-1] = bs[utf8.RuneCount(bs)-i-1], bs[i]
	}

	return string(bs), nil
}
