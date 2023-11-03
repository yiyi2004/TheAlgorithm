package main

import "fmt"

func main() {
	n1 := new(ListNode)
	n2 := new(ListNode)
	n3 := new(ListNode)
	n4 := new(ListNode)
	n1.Val = 1
	n2.Val = 2
	n3.Val = 3
	n4.Val = 4
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	res := swapPairs(n1)
	printList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head.Next
	cur := head
	next := new(ListNode)
	for cur != nil {
		if cur.Next == nil {
			break
		}
		next = cur.Next.Next

		tmpA := cur
		tmpB := cur.Next
		if tmpB.Next == nil {
			tmpA.Next = nil
		} else if tmpB.Next.Next == nil {
			tmpA.Next = tmpB.Next
		} else {
			tmpA.Next = tmpB.Next.Next
		}
		tmpB.Next = tmpA

		cur = next
	}

	return res
}
