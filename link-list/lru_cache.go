package main

import "fmt"

// LRUCache 的 Get 操作很简单，在 map 中直接读取双向链表的结点。
// 如果 map 中存在，将它移动到双向链表的表头，并返回它的 value 值，
// 如果 map 中不存在，返回 -1。

// LRUCache 的 Put 操作也不难。
// 先查询 map 中是否存在 key，如果存在，更新它的 value，
// 并且把该结点移到双向链表的表头。如果 map 中不存在，
// 新建这个结点加入到双向链表和 map 中。
// 最后别忘记还需要维护双向链表的 cap，如果超过 cap，
// 需要淘汰最后一个结点，双向链表中删除这个结点，map 中删掉这个结点对应的 key。

type ListNode struct {
	Key, Val   int
	Prev, Next *ListNode
}

func newListNode(key int, val int) *ListNode {
	return &ListNode{
		Key:  key,
		Val:  val,
		Prev: nil,
		Next: nil,
	}
}

type LRUCache struct {
	head, tail *ListNode
	Keys       map[int]*ListNode
	Cap        int
}

// 新建 LRUCache
func Constructor(capacity int) *LRUCache {
	return &LRUCache{Keys: make(map[int]*ListNode), Cap: capacity}
}

// 访问数据
func (l *LRUCache) Get(key int) int {
	if node, ok := l.Keys[key]; ok {
		l.Remove(node)
		l.Add(node)
		return node.Val
	}
	return -1
}

// 插入数据
func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.Keys[key]; ok {
		node.Val = value
		l.Remove(node)
		l.Add(node)
		return
	} else {
		node = &ListNode{Key: key, Val: value}
		l.Keys[key] = node
		l.Add(node)
	}
	if len(l.Keys) > l.Cap {
		delete(l.Keys, l.tail.Key)
		l.Remove(l.tail)
	}
}

// 添加节点 头插
func (l *LRUCache) Add(node *ListNode) {
	node.Prev = nil
	node.Next = l.head
	if l.head != nil {
		l.head.Prev = node
	}
	l.head = node
	if l.tail == nil {
		l.tail = node
		l.tail.Next = nil
	}
}

// 删除节点
func (l *LRUCache) Remove(node *ListNode) {
	if node == l.head {
		l.head = node.Next
		if node.Next != nil {
			node.Next.Prev = nil
		}
		node.Next = nil
		return
	}
	if node == l.tail {
		l.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *LRUCache) String() string {
	p := l.tail
	var res string = ""
	for p != nil {
		res += fmt.Sprintf("%v:%v - ", p.Key, p.Val)
		p = p.Prev
	}
	res += "\n"
	return res
}

func main() {
	fmt.Println("vim-go")
	//l := NewCache(10)
	l := Constructor(10)
	l.Put(1, 1)
	fmt.Println(l)

	for i := 0; i < 12; i++ {
		l.Put(i+1, i*2)
		fmt.Printf("%v", l)
		if i%3 == 0 {
			l.Get(i - 1)
			fmt.Printf("%v", l)
		}
	}
}
