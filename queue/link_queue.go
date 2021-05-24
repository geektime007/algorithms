package main

import "fmt"

type Node struct {
	Val  string
	Next *Node
}

func NewNode(val string) *Node {
	return &Node{
		Val:  val,
		Next: nil,
	}
}

type LinkQueue struct {
	// 队列容量
	length int
	// 队列当前长度
	count int
	// 队头
	head *Node
	// 对尾
	tail *Node
}

// 初始化队列
func NewLinkQueue(n int) *LinkQueue {
	node := NewNode("")
	return &LinkQueue{
		length: n,
		count:  0,
		head:   node,
		tail:   node,
	}
}

// 入队
func (q *LinkQueue) EnQueue(val string) bool {
	if q == nil || q.tail == nil || q.count == q.length {
		return false
	}
	newNode := NewNode(val)
	q.tail.Next = newNode
	q.tail = newNode
	q.count += 1
	return true
}

// 出队
func (q *LinkQueue) DeQueue() string {
	if q == nil || q.head == nil || q.count == 0 || q.head == q.tail {
		return ""
	}
	if q.head.Next == q.tail {
		q.tail = q.head
	}
	data := q.head.Next.Val
	q.head.Next = q.head.Next.Next
	q.count -= 1
	return data
}

func (q *LinkQueue) String() string {
	var res string
	i := 0
	p := q.head.Next
	for p != nil {
		res += fmt.Sprintf("%v|%v - ", i, p.Val)
		p = p.Next
		i += 1
	}
	res += fmt.Sprintf("\n")
	return res
}

func main() {
	q := NewLinkQueue(15)
	for i := 0; i < 26; i++ {
		fmt.Printf("Enqueue: %v\n", q.EnQueue(string(int('a')+i)))
		fmt.Println(q)
		if i%5 == 0 {
			fmt.Printf("DeQueue: %v\n", q.DeQueue())
			fmt.Println(q)
		}
		fmt.Println("++++++++++++++")
	}

}
