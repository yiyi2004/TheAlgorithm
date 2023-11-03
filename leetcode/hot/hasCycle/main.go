package main

import "fmt"

func main() {
	n1 := &ListNode{}
	n2 := &ListNode{}
	n3 := &ListNode{}
	n1.Val = 1
	n2.Val = 2
	n3.Val = 3
	n1.Next = n2
	n2.Next = n3
	n3.Next = n2
	res := detectCycle(n1)
	fmt.Println(res.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//func hasCycle(head *ListNode) bool {
//	m := make(map[*ListNode]struct{})
//	for head != nil {
//		if exists(m, head) {
//			return true
//		}
//		m[head] = struct{}{}
//		head = head.Next
//	}
//	return false
//}
//
//func exists(m map[*ListNode]struct{}, val *ListNode) bool {
//	_, ok := m[val]
//	return ok
//}

// 通过快慢指针解决问题，快指针总是能够追上慢指针
//func hasCycle(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//		return false
//	}
//	slow := head
//	fast := head.Next
//	for slow != fast {
//		if slow == nil || fast == nil || fast.Next == nil {
//			return false
//		}
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return true
//}
//
//func detectCycle(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return nil
//	}
//	slow := head
//	fast := head.Next
//	for slow != fast {
//		if slow == nil || fast == nil || fast.Next == nil {
//			return nil
//		}
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	return slow
//}
//
//func detectCycle1(head *ListNode) *ListNode {
//	slow, fast := head, head
//	for fast != nil {
//		slow = slow.Next
//		if fast.Next == nil {
//			return nil
//		}
//		fast = fast.Next.Next
//		if fast == slow {
//			p := head
//			for p != slow {
//				p = p.Next
//				slow = slow.Next
//			}
//			return p
//		}
//	}
//	return nil
//}

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if exists(m, head) {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}
func exists(m map[*ListNode]struct{}, val *ListNode) bool {
	_, ok := m[val]
	return ok
}
