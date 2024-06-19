package main

import "fmt"

func main() {
	fmt.Println(letterCombinations(""))
}

var lettersByNumber map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func backtrack(combination *[]string, digits string, index int, path string) {
	if len(path) == len(digits) {
		*combination = append(*combination, path)
		return
	}

	possibleLetters := lettersByNumber[string(digits[index])]
	for _, possibleLetter := range possibleLetters {
		backtrack(combination, digits, index+1, path+string(possibleLetter))
	}
}

func letterCombinations(digits string) []string {
	var combination []string

	backtrack(&combination, digits, 0, "")

	return combination
}
