package main

import "fmt"

type Stack struct {
	items  []rune
	count  int
	length int
}

func NewStack(n int) *Stack {
	return &Stack{
		items:  make([]rune, n),
		count:  0,
		length: n,
	}
}

func (s *Stack) Push(item rune) bool {
	if s.count == s.length {
		return false
	}
	s.items[s.count] = item
	s.count++
	return true
}

func (s *Stack) Pop() rune {
	if s.count == 0 {
		return rune(0)
	}
	tmp := s.items[s.count-1]
	s.count--
	return tmp
}

func (s *Stack) IsEmpty() bool {
	if s.count <= 0 {
		return true
	}
	return false
}

func (s *Stack) Get() rune {
	if s.count == 0 {
		return rune(0)
	}
	fmt.Println("===", s.items[s.count-1])
	return s.items[s.count-1]
}

func (s *Stack) Println() {
	//fmt.Println(s.count)
	for _, v := range s.items {
		fmt.Printf("%v-", v)
	}
	fmt.Printf("\n")
}

func isValid(s string) bool {
	if s == "" {
		return true
	}
	stack := NewStack(len(s))
	for i, vv := range s {
		if vv == rune(0) {
			break
		}
		if stack.IsEmpty() {
			stack.Push(rune(s[i]))
			stack.Println()
		} else {
			v := stack.Get()
			fmt.Printf("Get=> %v %T %v %T\n",
				v, v, rune(s[i]), rune(s[i]))
			//v, v, rune(vv), rune(vv))
			if v == rune(s[i]) {
				fmt.Println("Pop => ", stack.Pop())
			} else {
				stack.Push(rune(s[i]))
				stack.Println()
			}
		}
	}
	return stack.IsEmpty()
}

func main() {
	fmt.Println("vim-go")
	var s string

	s = "()"
	fmt.Println(s, " => ", isValid(s))

	s = "()[]{}"
	fmt.Println(s, " => ", isValid(s))

	s = "(]"
	fmt.Println(s, " => ", isValid(s))

	s = "({)]"
	fmt.Println(s, " => ", isValid(s))

	s = "{[]}"
	fmt.Println(s, " => ", isValid(s))
}
