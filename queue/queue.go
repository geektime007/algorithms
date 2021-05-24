package main

import "fmt"

type Queue struct {
	// 数组：items 数组大小 n
	items  []string
	length int
	// head 表示队头下标
	head int
	// tail 表示队尾下标
	tail int
}

// 申请一个大小为 n 的数组
func NewArrayQueue(n int) *Queue {
	return &Queue{
		items:  make([]string, n),
		length: n,
		head:   0,
		tail:   0,
	}
}

// 入队
func (q *Queue) EnQueue(item string) bool {
	// 如果 tail == n 表示队列已经满了
	if q.tail == q.length {
		return false
	}
	q.items[q.tail] = item
	q.tail++
	return true
}

// 出队
func (q *Queue) DeQueue() string {
	// 如果 head == tail 表示队列为空
	if q.head == q.tail {
		return ""
	}
	ret := q.items[q.head]
	q.head++
	return ret
}

func main() {
	fmt.Println("vim-go")
}
