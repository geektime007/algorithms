package main

import "fmt"

func reverse(s []byte) []byte {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

func reverse2(s string) string {
	str := []rune(s)
	for i, j := 0, len(s)-1; i < j; i++ {
		str[i], str[j] = str[j], str[i]
		j--
	}
	return string(str)
}

func reverse1(s string) string {
	str := []rune(s)
	left := 0
	right := len(str) - 1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	return string(str)
}

func main() {
	fmt.Println("vim-go")
	s := "abc"
	fmt.Println(s)
	s1 := reverse([]byte(s))
	fmt.Println(string(s1))
	s2 := reverse1(s)
	fmt.Println(string(s2))

	ss := "a你好b 世界c"
	fmt.Println(ss)
	ss1 := reverse1(ss)
	fmt.Println(string(ss1))

}
