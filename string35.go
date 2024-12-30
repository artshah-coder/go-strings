// Given strings S and S0. Remove from the string S all substrings that match S0.
// If there are no matching substrings, then print the string S without changes.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f := os.Stdin
	defer f.Close()

	fmt.Print("The programm removes all occurrences the sting 2 in string 1.\n",
		"Please, input stings.\nString 1: ")
	s1, s2 := "", ""
	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		fmt.Println("An error while reading data!")
		return
	}
	s1 = scanner.Text()
	fmt.Print("String 2: ")
	if !scanner.Scan() {
		fmt.Println("An error while reading data!")
		return
	}
	s2 = scanner.Text()
	fmt.Print("You have inputed \"", s1, "\" as string 1, \"", s2, "\" as string 2.\n")
	result := ""
	before, after, found := strings.Cut(s1, s2)
	if found {
		for {
			s1 = after
			result += before
			if before, after, found = strings.Cut(s1, s2); !found {
				result += s1
				break
			}
		}
	} else {
		result = s1
	}
	fmt.Println("Result:", result)

}
