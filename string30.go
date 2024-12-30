// Given a symbol C and strings S, S0.
// After each occurrence of the character C in the string S, insert the string S0
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	fmt.Println("The programm finds occuring symbols-separator in the first string",
		"and pastes the second string in the first string after them")
	fmt.Println("Input 2 strings and symbol-separator:")
	s1, s2 := "", ""
	var sep rune
	scanner := bufio.NewScanner(f)
	fmt.Print("the first string: ")
	if !scanner.Scan() {
		fmt.Println("An error while reading data!")
		return
	} else {
		s1 = scanner.Text()
	}
	fmt.Print("the second string: ")
	if !scanner.Scan() {
		fmt.Println("An error while reading data!")
		return
	} else {
		s2 = scanner.Text()
	}
	fmt.Print("symbol-separator: ")
	if _, err := fmt.Scanf("%c", &sep); err != nil {
		fmt.Println("An error while reading data!")
		return
	}
	fmt.Println("You have inputed ", s1, s2, string(sep))

	strs := []string{}
	for before, after, found := strings.Cut(s1, string(sep)); ; {
		strs = append(strs, before)
		strs = append(strs, string(sep))
		strs = append(strs, s2)
		s1 = after
		if before, after, found = strings.Cut(s1, string(sep)); !found {
			strs = append(strs, s1)
			break
		}
	}
	fmt.Println("Result:", strings.Join(strs, ""))
}
