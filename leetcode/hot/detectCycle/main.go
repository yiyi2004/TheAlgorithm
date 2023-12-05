package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next
	for fast != nil {
		if fast == slow {
			return slow
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
	return nil
}
