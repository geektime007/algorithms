package main

import "fmt"

// 字符串替换空格
func replaceSpace(s string) string {
	str := ""
	for _, v := range s {
		if v == ' ' {
			str += "%20"
		} else {
			str += string(v)
		}
	}
	return str
}

func main() {
	a := "a 你好 b 世界 c"
	fmt.Println(a)
	b := replaceSpace(a)
	fmt.Println(b)
}
