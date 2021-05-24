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

func (q *Queue) EnQueue1(item string) bool {
	// tail == n 表示队列末尾没有空间了
	if q.tail == q.length {
		// tail == n && head == 0, 表示整个队列都占满了
		if q.head == 0 {
			return false
		}
		// 数据搬移
		for i := q.head; i < q.tail; i++ {
			q.items[i-q.head] = q.items[i]
		}
		// 搬移完之后重新更新 head 和 tail
		q.tail -= q.head
		q.head = 0
	}
	q.items[q.tail] = item
	q.tail += 1
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

func (q *Queue) String() string {
	var res string
	for i, v := range q.items {
		res += fmt.Sprintf("%v|%v - ", i, v)
	}
	res += fmt.Sprintf("\n")
	return res
}

func main() {
	fmt.Println("vim-go")

	q := NewArrayQueue(15)
	for i := 0; i < 26; i++ {
		fmt.Printf("Enqueue: %v\n", q.EnQueue1(string(int('a')+i)))
		fmt.Println(q)
		if i%2 == 0 {
			fmt.Printf("DeQueue: %v\n", q.DeQueue())
			fmt.Println(q)
		}
		fmt.Println("++++++++++++++")
	}
}
