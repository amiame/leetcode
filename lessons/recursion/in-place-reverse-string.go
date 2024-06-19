package main

import "fmt"

func main() {
	input := []string{"h", "e", "l", "l", "o"}
	reverseString(input)
}

func reverseString(s []string) {
	if len(s) <= 0 {
		fmt.Printf("\n")
		return
	}

	fmt.Printf("%s", s[len(s)-1])

	reverseString(s[:len(s)-1])
}
