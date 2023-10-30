package main

import (
	"fmt"
)

func main() {
	n1 := new(ListNode)
	n2 := new(ListNode)
	n3 := new(ListNode)
	n4 := new(ListNode)
	n1.Val = 1
	n2.Val = 2
	n3.Val = 2
	n4.Val = 1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	palindrome := isPalindrome(n1)
	fmt.Println(palindrome)
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

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	length := getLen(head)
	if length == 1 {
		return true
	}

	var headA, headB *ListNode
	half := length >> 1
	if isOdd(length) {
		headA = head
		tmp := head
		tmpVal := 1
		for tmp != nil {
			if tmpVal == half {
				break
			}
			tmp = tmp.Next
			tmpVal++
		}
		headB = tmp.Next.Next
		tmp.Next = nil
	} else {
		headA = head
		tmp := head
		tmpVal := 1
		for tmp != nil {
			if tmpVal == half {
				break
			}
			tmp = tmp.Next
			tmpVal++
		}
		headB = tmp.Next
	}

	headB = reverseList(headB)
	flag := true
	for headA != nil && headB != nil {
		if headA.Val != headB.Val {
			flag = false
			break
		}
		headA = headA.Next
		headB = headB.Next
	}
	return flag
}

func isOdd(val int) bool {
	if val%2 == 0 {
		return false
	}
	return true
}

func getLen(head *ListNode) int {
	l := 0
	tmp := head
	for tmp != nil {
		l++
		tmp = tmp.Next
	}
	return l
}

func reverseList(head *ListNode) *ListNode {
	newHead := new(ListNode)
	for head != nil {
		next := head.Next
		head.Next = newHead.Next
		newHead.Next = head
		head = next
	}

	return newHead.Next
}
