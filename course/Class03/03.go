package Class03

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func NewNode(v int) *Node {
	return &Node{
		Value: v,
	}
}

type DoubleNode struct {
	Value int
	Pre   *DoubleNode
	Next  *DoubleNode
}

func NewDoubleNode(v int) *DoubleNode {
	return &DoubleNode{
		Value: v,
	}
}

func ReverseLinkedList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var pre, next *Node
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func ReverseDoubleLinkedList(head *DoubleNode) *DoubleNode {
	if head == nil {
		return nil
	}

	var pre, next *DoubleNode
	for head != nil {
		next = head.Next
		head.Next = pre
		head.Pre = next
		pre = head
		head = next
	}

	return pre
}

func printDoubleLinkedList(head *DoubleNode) {
	if head == nil {
		fmt.Println("head is nil")
		return
	}

	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func printLinkedList(head *Node) {
	if head == nil {
		fmt.Println("head is nil")
		return
	}

	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func removeValue(head *Node, num int) *Node {
	if head == nil {
		return nil
	}
	if head.Value == num {
		head = head.Next
		return head
	}

	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Value == num {
			pre.Next = cur.Next
			return head
		}
		pre = cur
		cur = cur.Next
	}
	return head
}

type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)

	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
	String() string
}

type RingArray struct {
	arr   []int
	size  int
	pushi int
	popi  int
	limit int
}

func NewRingArray(limit int) *RingArray {
	return &RingArray{
		arr:   make([]int, limit),
		limit: limit,
	}
}

func (ra *RingArray) Push(val int) error {
	if ra.size == ra.limit {
		return errors.New("array is full")
	}
	ra.arr[ra.pushi] = val
	ra.size++
	ra.pushi = next(ra.pushi, ra.limit)
	return nil
}

func (ra *RingArray) Pop() int {
	ra.size--
	val := ra.arr[ra.popi]
	ra.popi = next(ra.popi, ra.limit)
	return val
}

func (ra *RingArray) IsEmpty() bool {
	return ra.size == 0
}

func next(i int, limit int) int {
	if i < limit-1 {
		return i + 1
	}
	return 0
}

type MinStack struct {
	stack    Stack
	minStack Stack
}

// 暂时不想写了，有点烦

func NewMinStack(limit int) *MinStack {
	return &MinStack{}
}
