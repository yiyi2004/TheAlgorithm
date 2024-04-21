package main

func main() {

}

type LinkNode struct {
	val  int
	next *LinkNode
}

func mergeSort(head *LinkNode, tail *LinkNode) *LinkNode {
	// 归并排序思路
	if head == nil {
		return nil
	}
	if head.next == tail {
		return head
	}
	// 找到中间节点
	fast, slow := head, head
	for fast != nil {
		slow = slow.next
		if fast.next != nil {
			fast = fast.next.next
		} else {
			fast = nil
		}
	}
	rHead := slow.next
	slow.next = nil
	// head --> slow
	// slow.next --> nil
	// 排序左边
	lLink := mergeSort(head, slow)
	// 排序右边
	rLink := mergeSort(rHead, nil)
	// 合并链表
	return merge(lLink, rLink)
}

func merge(lHead, rHead *LinkNode) *LinkNode {
	return nil
}
