package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	InsertNodeFromHead(n *ListNode)
	// 从链表尾部插入一个元素
	InsertFromTail(val int)
	// 从链表尾部插入一个节点
	InsertNodeFromTail(n *ListNode)
	// 链表反转
	// Bad method
	Reversed() *ListNode
	// 链表反转1
	Reversed1() *ListNode
	// 链表反转2
	Reversed2() *ListNode
	// 链表反转前 N 个节点
	Reversed3(n int) *ListNode
	// 链表反转递归实现
	Reversed4() *ListNode
	// 从链表第 index 个节点开始反转
	ReversedFromIndex(index int)
	// 从链表第 start 个节点开始反转, 到 end 个节点停止反转
	ReversedFromStart2End(start int, end int) *ListNode
	// 获取链表长度
	GetLength() int
	// 查找元素位置
	Search(val int) int
	// 获取指定index位置的元素
	GetVal(index int) int
	// 获取所有的元素
	GetAll() []int
	// 检查单链表中是否存在环
	HasCycle(l *ListNode) bool
}

// 新建一个节点
func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// 后面添加
func (head *ListNode) Add(val int) {
	point := head
	for point.Next != nil {
		point = point.Next
	}
	node := NewNode(val)
	point.Next = node
}

// 删除指定 index  位置元素
func (head *ListNode) Delete(index int) int {
	if index < 0 || index > head.GetLength() {
		return -1
	}
	point := head
	for i := 0; i < index; i++ {
		point = point.Next
	}
	data := point.Next.Val
	point.Next = point.Next.Next
	return data
}

// 在指定 index 位置插入元素
func (head *ListNode) Insert(index int, val int) {
	if index < 0 || index > head.GetLength() {
		fmt.Println("insert value failed")
		return
	}
	point := head
	for i := 0; i < index-1; i++ {
		point = point.Next
	}
	newNode := &ListNode{Val: val}
	newNode.Next = point.Next
	point.Next = newNode
}

// 从链表头部插入一个元素
func (head *ListNode) InsertFromHead(val int) {
	newNode := NewNode(val)
	if head.GetLength() == 1 {
		head.Next = newNode
		return
	}
	newNode.Next = head.Next
	head.Next = newNode
}

// 从链表头部插入一个节点
func (head *ListNode) InsertNodeFromHead(n *ListNode) {
	if head.GetLength() == 0 {
		head.Next = n
		return
	}
	n.Next = head.Next
	head.Next = n
}

// 从链表尾部插入一个元素
func (head *ListNode) InsertFromTail(val int) {
	newNode := NewNode(val)
	point := head
	for point.Next != nil {
		point = point.Next
	}
	point.Next = newNode
}

// 从链表尾部插入一个节点
func (head *ListNode) InsertNodeFromTail(n *ListNode) {
	point := head
	for point.Next != nil {
		point = point.Next
	}
	point.Next = n
}

// 链表反转
// Bad method
func (head *ListNode) Reversed() *ListNode {
	point := head.Next
	var res *ListNode = NewNode(0)
	for point.Next != nil {
		res.InsertNodeFromHead(NewNode(point.Val))
		point = point.Next
	}
	res.InsertNodeFromHead(NewNode(point.Val))
	return res
}

// 链表反转1
func (head *ListNode) Reversed1() *ListNode {
	cur := head.Next
	var prev *ListNode = nil
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	head.Next = prev
	return head
}

// 链表反转2
func (head *ListNode) Reversed2() *ListNode {
	cur := head.Next
	var prev *ListNode = nil
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	head.Next = prev
	return head
}

// 递归内部实现
func (head *ListNode) _reversed4(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	newHead := head._reversed4(l.Next)
	l.Next.Next = l
	l.Next = nil
	return newHead
}

// 链表反转4 递归实现
func (head *ListNode) Reversed4() *ListNode {
	head.Next = head._reversed4(head.Next)
	return head
}

// 链表反转前 N 个节点
func (head *ListNode) Reversed3(n int) *ListNode {
	cur := head.Next
	var prev *ListNode = nil
	count := 0
	for cur != nil {
		count += 1
		if count == n {
			// 这里将当前 cur 链, 直接插到 prev 尾部即可
			prev.InsertNodeFromTail(cur)
			break
		}

		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	head.Next = prev
	return head
}

// 从链表第 index 个节点开始反转
func (head *ListNode) ReversedFromIndex(index int) *ListNode {
	length := head.GetLength()
	if index < 0 || index > length {
		return nil
	}
	cur := head.Next
	for i := 0; i < length; i++ {
		if i < index-1 {
			cur = cur.Next
			continue
		}
		if i == index {
			cur.Reversed2()
			break
		}
	}
	return head
}

// 从链表第 start 个节点开始反转, 到 end 个节点停止反转
func (head *ListNode) ReversedFromStart2End(start int, end int) *ListNode {
	length := head.GetLength()
	if start < 0 || start > end || end > length {
		return nil
	}
	cur := head.Next
	for i := 0; i < length; i++ {
		if i < start-1 {
			cur = cur.Next
			continue
		}
		if i == start {
			cur.Reversed3(end - start)
			break
		}
	}
	return head
}

// 获取链表长度
func (head *ListNode) GetLength() int {
	length := 0
	point := head
	for point.Next != nil {
		point = point.Next
		length += 1
	}
	return length
}

// 查找元素位置
func (head *ListNode) Search(val int) int {
	point := head.Next
	index := 0
	for point.Next != nil {
		if point.Val == val {
			return index
		} else {
			point = point.Next
			index += 1
		}
	}
	if point.Val == val {
		return index
	}
	return -1
}

// 获取指定index位置的元素
func (head *ListNode) GetVal(index int) int {
	if index < 0 || index > head.GetLength() {
		return -1
	}
	point := head.Next
	for i := 0; i < index-1; i++ {
		point = point.Next
	}
	return point.Val
}

// 获取所有的元素
func (head *ListNode) GetAll() []int {
	if head.Next == nil {
		return nil
	}
	var list []int
	point := head.Next
	for point.Next != nil {
		list = append(list, point.Val)
		point = point.Next
	}
	list = append(list, point.Val)
	return list
}

// 初始化数组为链表
func InitLinkNodes(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := NewNode(0)
	fmt.Println("初始化链表:", values)
	for i, _ := range values {
		head.Add(values[i])
	}
	return head
}

// 打印链表
func PrintLinkNodes(l *ListNode) {
	if l == nil {
		panic("list is nil")
		return
	}
	point := l.Next
	fmt.Printf("打印链表：")
	for point.Next != nil {
		fmt.Printf("%v -> ", point.Val)
		point = point.Next
	}
	fmt.Printf("%v -> nil\n", point.Val)
}

// 将两个链表合并成为一个有序链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val {
		res = l2
		res.Next = MergeTwoLists(l2.Next, l1)
	} else {
		res = l1
		res.Next = MergeTwoLists(l1.Next, l2)
	}
	return res
}

// 将两个链表合并成为一个有序链表  非递归方式
func MergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := NewNode(0)
	prev := head
	var cur *ListNode
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

// 检查单链表中是否存在环
func (head *ListNode) HasCycle() bool {
	cur := head.Next
	for cur != nil {
		p := cur.Next
		// 这种方法不对,  有可能会造成死循环
		for p != nil {
			if cur == p {
				return true
			}
			p = p.Next
		}
		cur = cur.Next
	}
	return false
}

// 检查单链表中是否存在环
// 标准答案  快慢指针
func (head *ListNode) HasCycle1() bool {
	slow, fast := head.Next, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 检查单链表中是否存在环
// 标准答案  快慢指针
func (head *ListNode) HasCycle2() (bool, *ListNode) {
	slow, fast := head.Next, head.Next
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true, slow
		}
	}
	return false, nil
}

// 检查单链表中是否存在环
func HasCycle(l *ListNode) bool {
	return l.HasCycle1()
}

// 检查单链表中是否存在环
// 并返回链表存在环的起点指针
func DetectCycle(l *ListNode) *ListNode {
	if l == nil || l.Next == nil {
		return nil
	}
	isCycle, slow := l.HasCycle2()
	if !isCycle {
		return nil
	}

	fast := l.Next
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// 创建一个环形链表
func CreateCycleList(values []int) *ListNode {
	l := InitLinkNodes(values)
	cur := l.Next
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = l
	return l
}

func main() {
	fmt.Println("vim-go")

	var nodes1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var nodes2 = []int{1, 3, 4}

	// 链表初始化
	l1 := InitLinkNodes(nodes1)
	l2 := InitLinkNodes(nodes2)
	PrintLinkNodes(l1)
	PrintLinkNodes(l2)

	l3 := NewNode(0)
	fmt.Println("合并链表：")
	l3.Next = MergeTwoLists(l1.Next, l2.Next)
	PrintLinkNodes(l3)

	fmt.Println("反转链表：")
	l4 := l1.Reversed2()
	PrintLinkNodes(l4)

	fmt.Println("删除链表节点：")
	l4.Delete(1)
	PrintLinkNodes(l4)

	fmt.Println("插入链表节点：")
	l4.Insert(2, 2)
	PrintLinkNodes(l4)

	// 链表插入
	l5 := InitLinkNodes(nodes1)
	PrintLinkNodes(l5)
	fmt.Println("反转链表节点：")
	PrintLinkNodes(l5.Reversed())

	fmt.Println("获取所有列表节点：")
	fmt.Println(l5.GetAll())
	fmt.Println("查找列表节点：")
	fmt.Println(l5.Search(2))
	fmt.Println("尾部插入节点：")
	l5.InsertFromTail(6)
	PrintLinkNodes(l5)
	fmt.Println("尾部插入节点：")
	l5.InsertNodeFromTail(NewNode(7))
	PrintLinkNodes(l5)
	fmt.Println("头部插入节点：")
	l5.InsertFromHead(3)
	PrintLinkNodes(l5)

	// 链表反转
	l6 := InitLinkNodes(nodes1)
	PrintLinkNodes(l6)

	fmt.Println("反转链表节点：")
	l6.Reversed1()
	PrintLinkNodes(l6)
	l6.Reversed2()
	PrintLinkNodes(l6)
	l6.Reversed4()
	PrintLinkNodes(l6)

	// 将链表从第 index 位开始反转

	fmt.Println("将链表从第 index 位开始反转：")
	l7 := l6.ReversedFromIndex(6)
	PrintLinkNodes(l7)
	PrintLinkNodes(l6)

	l8 := InitLinkNodes(nodes1)
	PrintLinkNodes(l8)
	// 将链表第 m 位到 n 位反转
	fmt.Println("将链表第 m 位到 n 位反转：")
	l8.ReversedFromStart2End(2, 6)
	PrintLinkNodes(l8)

	// 初始化一个环形链表
	fmt.Println("初始化一个环形链表：")
	l9 := CreateCycleList(nodes1)
	fmt.Println("是否存在环: ", l9.HasCycle())

	fmt.Println("环的起始节点: ", DetectCycle(l9))

}
