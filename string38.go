// Given strings S, S1 and S2. Replace all occurrences of string S1 in string S with string S2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s1, s2, s3 := "", "", ""
	f := os.Stdin
	defer f.Close()

	fmt.Println("The programm replaces all substring 2 occurrences in string 1 to substring 3.")
	fmt.Println("Please, input 3 strings.")
	fmt.Print("String 1: ")
	scanner := bufio.NewScanner(f)
	ok := scanner.Scan()
	if ok == false {
		fmt.Println("Error reading data!")
		return
	}
	s1 = scanner.Text()
	fmt.Print("Sting 2: ")
	ok = scanner.Scan()
	if ok == false {
		fmt.Println("Error reading data!")
		return
	}
	s2 = scanner.Text()
	fmt.Print("String 3: ")
	ok = scanner.Scan()
	if ok == false {
		fmt.Println("Error reading data!")
		return
	}
	s3 = scanner.Text()

	fmt.Println("Result:", strings.ReplaceAll(s1, s2, s3))
}
