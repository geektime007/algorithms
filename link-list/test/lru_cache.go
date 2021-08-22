package main

import "fmt"

type ListNode struct {
	Prev, Next *ListNode
	Key, Val   int
}

func NewListNode(key int, val int) *ListNode {
	return &ListNode{Key: key, Val: val}
}

type LRUCache struct {
	head, tail *ListNode
	Keys       map[int]*ListNode
	Cap        int
}

func NewLRUCache(Cap int) *LRUCache {
	return &LRUCache{
		Keys: make(map[int]*ListNode, 0),
		Cap:  Cap,
	}
}

func (this *LRUCache) String() string {
	cur := this.tail
	res := ""
	for cur != nil {
		res += fmt.Sprintf("%v|%v -", cur.Key, cur.Val)
		cur = cur.Prev
	}
	res += "\n"
	return res
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, val int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = val
		this.Remove(node)
		this.Add(node)
		return
	}
	n := NewListNode(key, val)
	this.Keys[key] = n
	this.Add(n)

	if this.Cap <= len(this.Keys) {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *ListNode) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *ListNode) {
	if this.head == node {
		this.head = node.Next
		if node.Next != nil {
			node.Next.Prev = nil
		}
		node.Next = nil
		return
	}
	if this.tail == node {
		this.tail = node.Prev
		this.tail.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func main() {
	fmt.Println("vim-go")
	cache := NewLRUCache(10)
}
