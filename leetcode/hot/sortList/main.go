package main

import "fmt"

func main() {
	// [4,2,1,3]
	n1 := new(ListNode)
	n2 := new(ListNode)
	n3 := new(ListNode)
	n4 := new(ListNode)
	n1.Val = 4
	n2.Val = 2
	n3.Val = 1
	n4.Val = 3
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	res := sortList(n1)
	//printList(res)
	_ = res
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 面试官：下一位！
//func sortList(head *ListNode) *ListNode {
//	valueArray := make([]int, 0)
//	tmp := head
//	for tmp != nil {
//		valueArray = append(valueArray, tmp.Val)
//		tmp = tmp.Next
//	}
//	sort.Ints(valueArray)
//	resHead := new(ListNode)
//	tmpResHead := resHead
//	for i := 0; i < len(valueArray); i++ {
//		n := new(ListNode)
//		n.Val = valueArray[i]
//		tmpResHead.Next = n
//		tmpResHead = tmpResHead.Next
//	}
//	return resHead.Next
//}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

// 遵循的策略是左闭右开
func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	fast, slow := head, head
	for fast != tail {
		fast = fast.Next
		slow = slow.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow

	return merge(sort(head, mid), sort(mid, tail))
}

func merge(h1, h2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	temp, temp1, temp2 := dummyHead, h1, h2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
			temp = temp.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
			temp = temp.Next
		}
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

//func merge(head1, head2 *ListNode) *ListNode {
//	dummyHead := &ListNode{}
//	temp, temp1, temp2 := dummyHead, head1, head2
//	for temp1 != nil && temp2 != nil {
//		if temp1.Val <= temp2.Val {
//			temp.Next = temp1
//			temp1 = temp1.Next
//		} else {
//			temp.Next = temp2
//			temp2 = temp2.Next
//		}
//		temp = temp.Next
//	}
//	if temp1 != nil {
//		temp.Next = temp1
//	} else if temp2 != nil {
//		temp.Next = temp2
//	}
//	return dummyHead.Next
//}
//
//func sort(head, tail *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//
//	if head.Next == tail {
//		head.Next = nil
//		return head
//	}
//
//	slow, fast := head, head
//	for fast != tail {
//		slow = slow.Next
//		fast = fast.Next
//		if fast != tail {
//			fast = fast.Next
//		}
//	}
//
//	mid := slow
//	return merge(sort(head, mid), sort(mid, tail))
//}
//
//func sortList(head *ListNode) *ListNode {
//	return sort(head, nil)
//}
