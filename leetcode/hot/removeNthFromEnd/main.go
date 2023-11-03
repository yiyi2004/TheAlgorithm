package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	tmp := head
	for tmp != nil {
		length++
		tmp = tmp.Next
	}

	if n == length {
		return head.Next
	}

	tmp = head
	index := 1
	for index != length-n {
		index++
		tmp = tmp.Next
	}

	tmp.Next = tmp.Next.Next

	return head
}
