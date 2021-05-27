package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// 前序遍历
// 根 --> 左 --> 右
func PreOrder(node *TreeNode) {
	if node != nil {
		fmt.Printf("%v ", node.Val)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历
// 左 --> 中 --> 右
func InfixOrder(node *TreeNode) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("%v ", node.Val)
		InfixOrder(node.Right)
	}
}

// 后序遍历
// 左--> 右 --> 中
func PostOrder(node *TreeNode) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("%v ", node.Val)
	}
}

// 先序遍历  非递归方式
func PrevOrderNoRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	s := NewStack()
	s.Push(node)
	current, _ := s.Pop().(*TreeNode)
	for current != nil {
		fmt.Printf("%v ", current.Val)
		if right := current.Right; right != nil {
			s.Push(right)
		}
		if left := current.Left; left != nil {
			s.Push(left)
		}
		current, _ = s.Pop().(*TreeNode)
	}
	fmt.Println()
}

// 中序遍历 非递归方式
func InOrderNoRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	s := NewStack()
	current := node
	for s.Size() > 0 || current != nil {
		if current != nil {
			s.Push(current)
			current = current.Left
		} else {
			current, _ = s.Pop().(*TreeNode)
			fmt.Printf("%v ", current.Val)
			current = current.Right
		}
	}
	fmt.Println()
}

// 后序遍历 非递归方式
func PostOrderNoRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	s := NewStack()
	current := node
	var pre *TreeNode
	for current != nil || s.Size() > 0 {
		for current != nil {
			s.Push(current)
			current = current.Left
		}
		current, _ = s.Top().(*TreeNode)
		if current.Right == nil || pre == current.Right {
			fmt.Printf("%v ", current.Val)
			pre = current
			s.Pop()
			current = nil
		} else {
			current = current.Right
		}
	}
	fmt.Println()
}

// 后序遍历2 非递归
func PostOrderNoRecursion2(node *TreeNode) {
	s1, s2 := NewStack(), NewStack()
	s1.Push(node)

	for s1.Size() > 0 {
		current, _ := s1.Pop().(*TreeNode)
		s2.Push(current)

		if current.Left != nil {
			s1.Push(current.Left)
		}

		if current.Right != nil {
			s1.Push(current.Right)
		}
	}

	for s2.Size() > 0 {
		current, _ := s2.Pop().(*TreeNode)
		fmt.Printf("%v  ", current.Val)
	}
}

// 层序遍历
func LevelOrder(node *TreeNode) {
	if node == nil {
		return
	}
	listQueue := NewQueue()
	listQueue.Push(node)
	for listQueue.Len() > 0 {
		current, _ := listQueue.Pop().(*TreeNode)
		if current.Left != nil {
			fmt.Printf(" %v --- (父%v)  \n", current.Left.Val, current.Val)
			listQueue.Push(current.Left)
		}
		if current.Right != nil {
			listQueue.Push(current.Right)
			fmt.Printf("  (父%v) --- %v \n", current.Val, current.Right.Val)
		}
	}
}

func main() {
	t := NewTreeNode(1)
	t.Left = NewTreeNode(3)
	t.Right = NewTreeNode(6)
	t.Left.Left = NewTreeNode(4)
	t.Left.Right = NewTreeNode(5)
	t.Left.Left.Left = NewTreeNode(7)
	PrevOrderNoRecursion(t)
	InOrderNoRecursion(t)
	PostOrder(t)
}
