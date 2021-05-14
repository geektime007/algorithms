package main

import "fmt"

// 方法一:中心扩展方法
func ExpandPalindrome(s string, left int, right int) (int, string) {
	length := len(s)
	palindromeLen := 0
	maxStr := ""
	for left >= 0 && right <= length-1 {
		if s[left] == s[right] {
			left -= 1
			right += 1
			palindromeLen += 1
		} else {
			//fmt.Println("left:", left, "right:", right)
			//maxStr = s[left+1 : right] // 此处有坑
			break
		}
	}
	maxStr = s[left+1 : right]
	return palindromeLen, maxStr
}

func longestPalindromeSubStr(s string) (int, string) {
	length := len(s)
	if length <= 1 {
		return 0, s
	}

	maxLength := 1
	maxStr := ""

	for i := 0; i < length; i++ {
		k1, str1 := ExpandPalindrome(s, i-1, i+1)
		k2, str2 := ExpandPalindrome(s, i, i+1)
		len1 := 2*k1 + 1
		len2 := 2 * k2

		//fmt.Println("str1:", str1, "len1:", len1,
		//	"str2:", str2, "len2:", len2,
		//	"maxLength:", maxLength,
		//	"maxStr:", maxStr)
		if len1 >= maxLength {
			maxLength = len1
			maxStr = str1
		}
		if len2 >= maxLength {
			maxLength = len2
			maxStr = str2
		}
	}
	if maxLength == 1 {
		return maxLength, s[:1]
	}
	return maxLength, maxStr
}

func main() {
	fmt.Println("vim-go")
	var s, subStr string
	var length int = 0

	s = "babad"
	length, subStr = longestPalindromeSubStr(s)
	fmt.Println(s, " longestPalindromeSubStr => ", length, subStr, "\n")

	s = "cbbd"
	length, subStr = longestPalindromeSubStr(s)
	fmt.Println(s, " longestPalindromeSubStr => ", length, subStr, "\n")

	s = "a"
	length, subStr = longestPalindromeSubStr(s)
	fmt.Println(s, " longestPalindromeSubStr => ", length, subStr, "\n")

	s = "ac"
	length, subStr = longestPalindromeSubStr(s)
	fmt.Println(s, " longestPalindromeSubStr => ", length, subStr, "\n")

	s = "bb"
	length, subStr = longestPalindromeSubStr(s)
	fmt.Println(s, " longestPalindromeSubStr => ", length, subStr, "\n")

	s = "ccc"
	length, subStr = longestPalindromeSubStr(s)
	fmt.Println(s, " longestPalindromeSubStr => ", length, subStr, "\n")

}
