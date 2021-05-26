package main

import "fmt"

type LinkNoder interface {
	// 后面添加
	Add(val int)
	// 删除指定 index  位置元素
	Delete(index int) int
	// 在指定 index 位置插入元素
	Insert(index int, val int)
	// 从链表头部插入一个元素
	InsertFromHead(val int)
	// 从链表头部插入一个节点
	InsertNodeFromHead(n *Node)
	// 从链表尾部插入一个元素
	InsertFromTail(val int)
	// 从链表尾部插入一个节点
	InsertNodeFromTail(n *Node)
	// 链表反转
	// Bad method
	Reversed() *Node
	// 链表反转1
	Reversed1() *Node
	// 链表反转2
	Reversed2() *Node
	// 链表反转前 N 个节点
	Reversed3(n int) *Node
	// 链表反转递归实现
	Reversed4() *Node
	// 从链表第 index 个节点开始反转
	ReversedFromIndex(index int)
	// 从链表第 start 个节点开始反转, 到 end 个节点停止反转
	ReversedFromStart2End(start int, end int) *Node
	// 获取链表长度
	GetLength() int
	// 查找元素位置
	Search(val int) int
	// 获取指定index位置的元素
	GetVal(index int) int
	// 获取所有的元素
	GetAll() []int
	// 检查单链表中是否存在环
	HasCycle(l *Node) bool
}

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val:  val,
		Next: nil,
	}
}

func InitLinkNodes(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	head := NewNode(0)
	fmt.Println("初始化链表：", nums)
	for i, _ := range nums {
		head.Add(nums[i])
	}
	return head
}

func (l *Node) String() string {
	res := ""
	cur := l.Next
	for cur != nil {
		res += fmt.Sprintf("%d-> ", cur.Val)
		cur = cur.Next
	}
	res += fmt.Sprintf("nil")
	return res
}

// 链表尾插
func (l *Node) Add(val int) {
	p := l
	for p.Next != nil {
		p = p.Next
	}
	node := NewNode(val)
	p.Next = node
}

// 获取链表长度
func (l *Node) GetLength() int {
	p := l
	length := 0
	for p.Next != nil {
		p = p.Next
		length++
	}
	return length
}

// 查找元素位置
func (l *Node) Search(val int) int {
	cur := l.Next
	index := 0
	for cur != nil {
		if cur.Val == val {
			return index
		}
		cur = cur.Next
		index++
	}
	return -1
}

// 在链表头部插入新元素
func (l *Node) InsertFromHead(val int) {
	node := NewNode(val)
	if l.GetLength() == 1 {
		l.Next = node
		return
	}
	node.Next = l.Next
	l.Next = node
}

// 在链表头部插入结点
func (l *Node) InsertNodeFromHead(n *Node) {
	if l.GetLength() == 0 {
		l.Next = n
		return
	}
	n.Next = l.Next
	l.Next = n
}

// 在链表尾部插入新元素
func (l *Node) InsertFromTail(val int) {
	p := l
	for p.Next != nil {
		p = p.Next
	}
	node := NewNode(val)
	p.Next = node
}

// 在链表尾部插入结点
func (l *Node) InsertNodeFromTail(node *Node) {
	p := l
	for p.Next != nil {
		p = p.Next
	}
	p.Next = node
}

// 在链表下标为 index 处插入元素
func (l *Node) InsertByIndex(index int, val int) {
	p := l
	count := 0
	for p.Next != nil {
		if count < index {
			p = p.Next
			count++
		} else {
			break
		}
	}
	node := NewNode(val)
	node.Next = p
}

// 删除链表下标为 index 的结点
func (l *Node) DeleteByIndex(index int) int {
	if index < 0 || index > l.GetLength() {
		return -1
	}
	p := l
	count := 0
	for p.Next != nil {
		if count < index {
			p = p.Next
			count++
		} else {
			break
		}
	}
	data := p.Next.Val
	p.Next = p.Next.Next
	return data
}

// 找出链表中间的结点
func (l *Node) Middle() *Node {
	if l == nil || l.Next == nil {
		return nil
	}
	slow := l.Next
	fast := l.Next
	var prev *Node
	for fast != nil {
		fast = fast.Next
		prev = slow
		slow = slow.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return prev
}

// 链表反转
func (l *Node) Reversed() *Node {
	cur := l.Next
	var prev *Node
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	l.Next = prev
	return l
}

// 链表反转1
func (l *Node) Reversed1() *Node {
	cur := l.Next
	var prev *Node
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	l.Next = prev
	return l
}

// 链表反转递归实现
func (l *Node) Reversed2() *Node {
	l.Next = l._reversed2(l.Next)
	return l
}

// 递归实现
func (l *Node) _reversed2(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head._reversed2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 链表反转3
// Bad method
func (l *Node) Reversed3() *Node {
	cur := l.Next
	var newNode *Node = NewNode(0)
	for cur != nil {
		newNode.InsertNodeFromHead(NewNode(cur.Val))
		cur = cur.Next
	}
	return newNode
}

// 链表反转前 N 个节点
func (l *Node) ReversedN(n int) *Node {
	if n < 0 {
		return l
	}
	cur := l.Next
	var prev *Node
	count := 0
	for cur != nil {
		if count == n {
			// 当循环到 N 时，直接将剩下的结点
			// 插到反转后的链表后边即可
			prev.InsertNodeFromTail(cur)
			break
		}
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp

		count++
	}
	l.Next = prev
	return l
}

// 从链表第 N 个结点开始反转
func (l *Node) ReversedFromN(n int) *Node {
	length := l.GetLength()
	if length <= 0 || n > length {
		return nil
	}
	cur := l.Next
	count := 0
	for cur != nil {
		count++
		if count == n {
			cur.Reversed()
			break
		}
		cur = cur.Next
	}
	return l
}

// 从链表第 M 个结点开始反转，到第 N 个结点停止反转
func (l *Node) ReversedFromStart2End(m int, n int) *Node {
	length := l.GetLength()
	if m < 0 || n > length || m >= n {
		return nil
	}
	cur := l.Next
	count := 0
	for cur != nil {
		count++
		if count == m {
			cur.ReversedN(n - m) // 反转
			break
		}
		cur = cur.Next
	}
	return l
}

// 检查链表是否存储在环
func (l *Node) HasCycle() (*Node, bool) {
	if l == nil || l.Next == nil {
		return nil, false
	}
	slow := l.Next
	fast := l.Next
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow, true
		}
	}
	return nil, false
}

// 判断链表是否是回文链表
func (l *Node) IsPalindromic() bool {
	n1 := l.Next
	// 找出中间结点
	// 这里如果是偶数，返回n/2-1
	// 如果是奇数 返回中间结点
	mid := l.Middle()
	n2 := mid.Reversed().Next // 从中间结点反转后的链表
	for n1 != nil && n2 != nil {
		if n1.Val == n2.Val {
			n1 = n1.Next
			n2 = n2.Next
		} else {
			return false
		}
	}
	return true
}

// 创建一个存在环的链表
func CreateCycleLinkList(nums []int) *Node {
	l := InitLinkNodes(nums)
	cur := l.Next
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = l
	return l
}

// 检查单链表中是否存在环
// 并返回链表存在环的起点指针
func DetectCycle(l *Node) *Node {
	if l == nil || l.Next == nil {
		return nil
	}

	slow, isExist := l.HasCycle()
	if !isExist {
		return nil
	}

	fast := l.Next
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
		//fmt.Println(slow.Val, fast.Val)
	}
	return slow
}

// 将两个链表合并为一个有序链表   非递归方式
func MergeTwoLinkList(l1 *Node, l2 *Node) *Node {
	var cur *Node
	var head *Node = NewNode(0)
	prev := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur = l1
			l1 = l1.Next
		} else {
			cur = l2
			l2 = l2.Next
		}
		prev.Next = cur
		prev = cur
	}

	if l1 != nil {
		prev.Next = l1
	}

	if l2 != nil {
		prev.Next = l2
	}
	return head.Next
}

// 将两个有序链表合并为一个有序链表 递归方式
func MergeTwoLinkList1(l1 *Node, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var res *Node
	if l1.Val <= l2.Val {
		res = l1
		res.Next = MergeTwoLinkList1(l1.Next, l2)
	} else {
		res = l2
		res.Next = MergeTwoLinkList1(l1, l2.Next)
	}
	return res
}

// 判断两个链表是否相交
func IsIntersection(l1 *Node, l2 *Node) *Node {
	l1l2 := l1.Next
	l2l1 := l2.Next
	for l1l2 != l2l1 {
		if l1l2 != nil {
			l1l2 = l1l2.Next
		} else {
			l1l2 = l2
		}

		if l2l1 != nil {
			l2l1 = l2l1.Next
		} else {
			l2l1 = l1
		}
	}
	// 没有相交返回nil
	return l1l2
}

func main() {
	nums1 := []int{2, 1, 6, 3, 9, 4}
	nums2 := []int{3, 6, 7, 0, 7, 2, 9}
	nums3 := []int{2, 1, 1, 2}
	nums4 := []int{2, 1, 3, 1, 2}
	l1 := InitLinkNodes(nums1)
	l2 := InitLinkNodes(nums2)
	l3 := InitLinkNodes(nums3)
	l4 := InitLinkNodes(nums4)
	fmt.Println("初始化链表：", l1)
	fmt.Println("链表中间结点：", l1.Middle().Val)
	fmt.Println("链表中间结点：", l2.Middle().Val)
	fmt.Println("反转链表：", l1.Reversed())
	fmt.Println("反转链表：", l1.Reversed1())
	fmt.Println("反转链表前N个结点：", l1.ReversedN(3))
	fmt.Println("反转链表后N个结点：", l1.ReversedFromN(3))
	fmt.Println("反转链表：", l1.Reversed1())
	fmt.Println("反转链表从M到N结点：", l1.ReversedFromStart2End(3, 5))
	fmt.Println("反转链表从M到N结点：", l1.ReversedFromStart2End(3, 5))
	fmt.Println("反转链表(递归实现)：", l1.Reversed2())
	//fmt.Println("合并链表(递归实现)：", MergeTwoLinkList1(l1.Next, l2.Next))
	fmt.Println("合并链表(非递归实现)：", MergeTwoLinkList(l1.Next, l2.Next))
	l22 := CreateCycleLinkList(nums2)
	_, isExist := l22.HasCycle()
	fmt.Println("是否存在环：", isExist)
	slow := DetectCycle(l22)
	fmt.Println("检测链表中是否存在环：", slow.Val)

	l1.InsertNodeFromTail(l3)
	l4.InsertNodeFromTail(l3)
	fmt.Println("判断两个链表是否相交：", IsIntersection(l1, l4))

	fmt.Println("是否是回文链表:", l1.IsPalindromic())
	fmt.Println("是否是回文链表:", l3.IsPalindromic())
}
