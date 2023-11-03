package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	if len(lists) == 0 {
		return nil
	}
	return mergeList(lists, 0, len(lists)-1)
}

func mergeList(lists []*ListNode, l, r int) *ListNode {
	if r-l == 0 {
		return lists[l]
	}
	mid := (l + r) >> 1
	leftList := mergeList(lists, l, mid)
	rightList := mergeList(lists, mid+1, r)
	return mergeTwoList(leftList, rightList)
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	}
	resHead := new(ListNode)
	if list1.Val < list2.Val {
		resHead.Next = list1
		list1 = list1.Next
	} else {
		resHead.Next = list2
		list2 = list2.Next
	}
	tmp := resHead.Next

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			tmp = tmp.Next
			list1 = list1.Next
		} else {
			tmp.Next = list2
			tmp = tmp.Next
			list2 = list2.Next
		}
	}

	for list1 != nil {
		tmp.Next = list1
		tmp = tmp.Next
		list1 = list1.Next
	}
	for list2 != nil {
		tmp.Next = list2
		tmp = tmp.Next
		list2 = list2.Next
	}

	return resHead.Next
}
