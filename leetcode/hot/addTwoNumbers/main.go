package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var flag bool
	res := createNode(-1)
	pre := res
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val
		if flag {
			tmp += 1
		}
		if tmp >= 10 {
			flag = true
		} else {
			flag = false
		}
		tmp = tmp % 10
		pre.Next = createNode(tmp)
		pre = pre.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		tmp := l1.Val
		if flag {
			tmp += 1
		}
		if tmp >= 10 {
			flag = true
		} else {
			flag = false
		}
		tmp = tmp % 10
		pre.Next = createNode(tmp)
		pre = pre.Next

		l1 = l1.Next
	}
	for l2 != nil {
		tmp := l2.Val
		if flag {
			tmp += 1
		}
		if tmp >= 10 {
			flag = true
		} else {
			flag = false
		}
		tmp = tmp % 10
		pre.Next = createNode(tmp)
		pre = pre.Next

		l2 = l2.Next
	}
	if flag {
		pre.Next = createNode(1)
	}

	return res.Next
}

func createNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}
