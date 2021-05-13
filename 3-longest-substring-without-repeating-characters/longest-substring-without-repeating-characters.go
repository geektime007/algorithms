package main

import "fmt"

//给定一个字符串，找出不含有重复字符的最长子串的长度。
//
//示例：
//
//给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。
//
//给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。
//
//给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。

//int [26] 用于字母 ‘a’ - ‘z’或 ‘A’ - ‘Z’
//int [128] 用于ASCII码
//int [256] 用于扩展ASCII码
func lengthOfLongestSubstring(s string) int {
	// 最大子串长度
	maxStrLen := 0
	//重复字符的上次出现的索引,上一个重复字符和当前重复字符的之间的子串即有效子串
	p := 0
	// 字符下标字典
	charData := make(map[int32]int, 128)
	for i, char := range s {
		j, ok := charData[char]
		if ok {
			if j > p {
				p = j
			}
		}

		fmt.Println("char=", string(char), "i=", i, "j=", j, "maxStrLen=", maxStrLen, "p=", p)
		if i+1-p > maxStrLen {
			maxStrLen = i + 1 - p
		}
		charData[char] = i + 1
	}
	return maxStrLen
}

func main() {
	fmt.Println("vim-go")
	l := lengthOfLongestSubstring("abcabcbb")
	fmt.Println("len=", l)
}
