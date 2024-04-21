package main

import "fmt"

func main() {
	head1 := NewLinkNode(1)
	node1 := NewLinkNode(4)
	node2 := NewLinkNode(5)
	head1.Next = node1
	node1.Next = node2

	head2 := NewLinkNode(2)
	node21 := NewLinkNode(6)
	head2.Next = node21

	head3 := NewLinkNode(1)
	node31 := NewLinkNode(3)
	node32 := NewLinkNode(4)
	head3.Next = node31
	node31.Next = node32

	input := []*LinkNode{head1, head2, head3}
	res := MergeLists(input)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

//合并 K 个升序链表
//给你一个链表数组，每个链表都已经按升序排列。
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//示例 1：
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//示例 2：
//输入：lists = []
//输出：[]
//示例 3：
//输入：lists = [[]]
//输出：[]

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func NewLinkNode(val int) *LinkNode {
	return &LinkNode{Val: val}
}

func MergeLists(lists []*LinkNode) *LinkNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeLists(lists, 0, len(lists)-1)
}

func mergeLists(lists []*LinkNode, l, r int) *LinkNode {
	if r-l == 0 {
		return lists[l]
	}
	mid := (l + r) >> 1
	lHead := mergeLists(lists, l, mid)
	rHead := mergeLists(lists, mid+1, r)
	return merge(lHead, rHead)
}

func merge(head1, head2 *LinkNode) *LinkNode {
	newHead := new(LinkNode)
	tmp := newHead
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			tmp.Next = head1
			tmp = tmp.Next
			head1 = head1.Next
		} else {
			tmp.Next = head2
			tmp = tmp.Next
			head2 = head2.Next
		}
	}
	for head1 != nil {
		tmp.Next = head1
		tmp = tmp.Next
		head1 = head1.Next
	}
	for head2 != nil {
		tmp.Next = head2
		tmp = tmp.Next
		head2 = head2.Next
	}

	return newHead.Next
}
