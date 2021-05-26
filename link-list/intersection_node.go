package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 不必记录链表的长度，两条链表分别走一遍
// 较短的链表先检测到Next为空，之后开始走较长的链表
// 较长的链表继续走，检测到Next为空时，走较短的链表
// 如果有交点，第二轮一定会找到交点

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA := headA
	pB := headB

	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}
