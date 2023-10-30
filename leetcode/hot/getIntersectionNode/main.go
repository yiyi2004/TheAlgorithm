package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1 := getLen(headA)
	l2 := getLen(headB)
	for l1 < l2 {
		l2--
		headB = headB.Next
	}
	for l2 < l1 {
		l1--
		headA = headA.Next
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headB
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func getLen(head *ListNode) int {
	tmp := head
	result := 0
	for tmp != nil {
		result++
		tmp = tmp.Next
	}
	return result
}
