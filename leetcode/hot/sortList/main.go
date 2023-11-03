package main

import "sort"

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	valueArray := make([]int, 0)
	tmp := head
	for tmp != nil {
		valueArray = append(valueArray, tmp.Val)
		tmp = tmp.Next
	}
	sort.Ints(valueArray)
	resHead := new(ListNode)
	tmpResHead := resHead
	for i := 0; i < len(valueArray); i++ {
		n := new(ListNode)
		n.Val = valueArray[i]
		tmpResHead.Next = n
		tmpResHead = tmpResHead.Next
	}
	return resHead.Next
}

//func sortList(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//
//	resHead := new(ListNode)
//	tmpHead := head
//
//	for tmpHead != nil {
//		minNode := tmpHead
//		cur := tmpHead
//		for cur != nil {
//
//		}
//		tmpHead = tmpHead.Next
//	}
//
//}
