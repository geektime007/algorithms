package main

import "fmt"

func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = -1
	i, j := 0, -1
	for i < len(s)-1 {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	return next
}

func KMPSearch(target string, text string) int {
	i, j := 0, 0
	next := getNext(target)
	fmt.Println(next)
	for j < len(text) {
		if i == len(target)-1 && target[i] == text[j] {
			return j - i
		}
		if j == -1 || target[i] == text[j] {
			i++
			j++
		} else {
			i = next[i]
		}
	}
	return -1
}

func main() {
	a := "ABABACDEFT"
	fmt.Println(a)
	index := KMPSearch("AC", a)
	fmt.Println("Index:", index)
}
