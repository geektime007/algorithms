package main

import "fmt"

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		letters := phoneMap[string(digits[index])]
		for i := 0; i < len(letters); i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

func main() {
	fmt.Println("vim-go")
	digits := "23"
	ret := letterCombinations(digits)
	fmt.Println(ret)
}
