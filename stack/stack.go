package main

import "fmt"

// 基于数组实现的顺序栈
type Stack struct {
	items  []*string // 数组
	count  int       // 栈中元素个数
	length int       // 栈的大小
}

// 初始化数组, 申请一个大小为 n 的数组空间
func NewArrayStack(n int) *Stack {
	return &Stack{
		items:  make([]*string, n),
		count:  0,
		length: n,
	}
}

// 入栈操作
func (s *Stack) Push(item string) bool {
	// 数组空间不够了，直接返回 false, 入栈失败
	if s.count == s.length {
		return false
	}
	// 将 item 放到下标为 count 的位置, 并且 count 加一
	s.items[s.count] = &item
	s.count++
	return true
}

// 出栈操作
func (s *Stack) Pop() string {
	// 栈为空，则直接返回 ""
	if s.count == 0 {
		return ""
	}
	tmp := s.items[s.count-1]
	s.count--
	return *tmp
}

func main() {
	fmt.Println("vim-go")
	stack := NewArrayStack(10)
	for i := 0; i < 11; i++ {
		v := fmt.Sprintf("stack %d", i)
		res := stack.Push(v)
		fmt.Printf("Push: %v => %v\n", v, res)
	}
	for i := 0; i < 11; i++ {
		res := stack.Pop()
		fmt.Println("Pop:", " => ", res)
	}
}
