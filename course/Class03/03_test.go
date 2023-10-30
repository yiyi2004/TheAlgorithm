package Class03

import (
	"fmt"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	head := NewNode(1)
	n2 := NewNode(4)
	n3 := NewNode(3)
	n4 := NewNode(2)
	head.Next = n2
	n2.Next = n3
	n3.Next = n4

	printLinkedList(head)
	res := ReverseLinkedList(head)
	printLinkedList(res)
}

func TestReverseDoubleLinkedList(t *testing.T) {
	head := NewDoubleNode(1)
	n2 := NewDoubleNode(2)
	n3 := NewDoubleNode(3)
	n4 := NewDoubleNode(4)
	head.Next = n2
	n2.Pre = head
	n2.Next = n3
	n3.Pre = n2
	n3.Next = n4
	n4.Pre = n3

	printDoubleLinkedList(head)
	res := ReverseDoubleLinkedList(head)
	printDoubleLinkedList(res)
}

func Test_removeValue(t *testing.T) {
	head := NewNode(1)
	n2 := NewNode(4)
	n3 := NewNode(3)
	n4 := NewNode(2)
	head.Next = n2
	n2.Next = n3
	n3.Next = n4

	printLinkedList(head)
	// 因为 Golang 中是值传递的
	res := removeValue(head, 4)
	printLinkedList(res)
}

func TestRingArray(t *testing.T) {
	ra := NewRingArray(10)
	ra.Push(10)
	ra.Push(20)
	ra.Push(30)
	ra.Push(40)

	fmt.Println(ra.Pop())
	fmt.Println(ra.Pop())
	fmt.Println(ra.Pop())
	fmt.Println(ra.Pop())
}
