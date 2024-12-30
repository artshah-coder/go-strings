// Given positive integers N1 and N2 and strings S1 and S2. Get a new string from these strings
// containing the first N1 characters of string S1 and the last N2 characters of string S2
// (in that order).
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

const N1 int = 5
const N2 int = 7

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	strs := [2]string{}

	fmt.Println("N1 = ", N1, "\nN2 = ", N2, "\nSubstrings concatenate programm:",
		"\nthe program concatenates 2 inputed substrings: the first N1 symbols from the first string",
		"+ the last N2 symbols from the second string.")
	scanner := bufio.NewScanner(f)
	for i := 0; i < len(strs); i++ {
		fmt.Print("Please, input string ", i+1, " and press Enter: ")
		scanner.Scan()
		strs[i] = scanner.Text()
	}

	// Output and handling entered strigs and result output
	fmt.Println("Strings entered:")
	var result string
	for i := 0; i < len(strs); i++ {
		fmt.Printf("string %d = %s\n", i+1, strs[i])
		if i == 0 {
			if utf8.RuneCountInString(strs[i]) > N1 {
				strs[i] = strs[i][:N1]
				result += strs[i]
			}
		}
		if i == len(strs)-1 {
			if utf8.RuneCountInString(strs[i]) > N2 {
				strs[i] = strs[i][len(strs[i])-N2:]
				result += strs[i]
			}
		}
	}
	fmt.Println("The result: ", result)
}
