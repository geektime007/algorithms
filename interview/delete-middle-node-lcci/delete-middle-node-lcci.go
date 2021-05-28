package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (head *ListNode) Add(val int) {
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = NewListNode(val)
}

func (head *ListNode) DeleteMiddleNode(node *ListNode) {
	node.Val = node.Next.Val
	tmp := node.Next
	node.Next = tmp.Next
	tmp = nil
}

func (head *ListNode) GetMiddleNode() *ListNode {
	cur := head
	return cur.Next.Next
}

func (head *ListNode) String() string {
	cur := head
	res := ""
	for cur.Next != nil {
		res += fmt.Sprintf("%v ->", cur.Val)
		cur = cur.Next
	}
	res += fmt.Sprintf("%v ->", cur.Val)
	res += "nil"
	return res
}

func InitLinkList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	node := NewListNode(0)
	for i, _ := range nums {
		node.Add(nums[i])
	}
	return node
}

func DeleteMiddleNode(node *ListNode) {
	node.Val = node.Next.Val
	tmp := node.Next
	node.Next = tmp.Next
	tmp = nil
}

func main() {
	fmt.Println("vim-go")
	a := []int{4, 5, 1, 9}
	l := InitLinkList(a)
	fmt.Println(l)
	l.DeleteMiddleNode(l.GetMiddleNode())
	fmt.Println(l)
}
