// Given strings S, S1 and S2. Replace the last occurrence of string S1 in string S with string S2
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

	fmt.Print("The program replace the last occurrence of string 2 in string 1 to string 3.\n",
		"Please, input 3 strings.\n", "String 1: ")
	s1, s2, s3 := "", "", ""
	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		fmt.Println("Error while reading data!")
		return
	}
	s1 = scanner.Text()
	fmt.Print("String 2: ")
	if !scanner.Scan() {
		fmt.Println("Error while reading data!")
		return
	}
	s2 = scanner.Text()
	fmt.Print("String 3: ")
	if !scanner.Scan() {
		fmt.Println("Error while reading data!")
		return
	}
	s3 = scanner.Text()
	fmt.Print("You have inputed \"", s1, "\" as s1, \"", s2, "\" as s2, \"", s3, "\" as s3.\n")
	result := ""
	before, after, found := strings.Cut(s1, s2)
	if found {
		for {
			result += before
			s1 = after
			if before, after, found = strings.Cut(s1, s2); found {
				result += s2
			} else {
				result += s3 + s1
				break
			}
		}
	} else {
		result = s1
	}
	fmt.Println("Result:", result)
}
