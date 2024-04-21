package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func canSorted(lists []*ListNode) []bool {
	res := make([]bool, len(lists))
	for _, list := range lists {
		res = append(res, canSortedSingleList(list))
	}
	return res
}

func canSortedSingleList(head *ListNode) bool {
	count := 0
	tmp := head
	for tmp.Next != nil {
		if tmp.Next.Val < tmp.Val {
			count++
		}
		tmp = tmp.Next
	}
	if count == 0 {
		return true
	} else if count == 1 && tmp.Val <= head.Val {
		return true
	} else {
		return false
	}
}
