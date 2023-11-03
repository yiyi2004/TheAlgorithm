package main

import "fmt"

func main() {
	//[[3,null],[3,0],[3,null]]
	n1 := new(Node)
	n2 := new(Node)
	n3 := new(Node)
	//n4 := new(Node)
	//n5 := new(Node)
	n1.Val = 3
	n2.Val = 3
	n3.Val = 3
	//n4.Val = 10
	//n5.Val = 1
	n1.Next = n2
	n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5

	n1.Random = nil
	n2.Random = n1
	n3.Random = nil
	//n4.Random = n3
	//n5.Random = n1

	res := copyRandomList(n1)
	fmt.Println(res)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// index:index
	// 你这样做的时间复杂度是 O(n^2)
	randomAssistMap := make(map[int]int)
	index := 0
	// 创建 map[*Node]index
	//
	// 完整链表深拷贝
	tmp := head
	for tmp != nil {
		randNode := tmp.Random
		count := -index
		tmpHead := head
		for randNode != nil && tmpHead != randNode {
			count++
			tmpHead = tmpHead.Next
		}
		if randNode == nil {
			randomAssistMap[index] = 0
		} else {
			randomAssistMap[index] = count
		}
		index++
		tmp = tmp.Next
	}

	indexNodeMap := make(map[int]*Node)
	tmpHead := head.Next
	res := new(Node)
	res.Val = head.Val
	indexNodeMap[0] = res
	tmpRes := res
	index = 1

	for tmpHead != nil {
		n := new(Node)
		n.Val = tmpHead.Val
		indexNodeMap[index] = n
		index++
		tmpRes.Next = n
		tmpRes = tmpRes.Next
		tmpHead = tmpHead.Next
	}

	index = 0
	tmpRes = res
	tmpHead = head
	for tmpRes != nil {
		if randomAssistMap[index] == 0 {
			if tmpHead.Random != nil {
				tmpRes.Random = tmpRes
			} else {
				tmpRes.Random = nil
			}
			index++
			tmpRes = tmpRes.Next
			tmpHead = tmpHead.Next
			continue
		}
		tmpRes.Random = indexNodeMap[randomAssistMap[index]+index]
		tmpRes = tmpRes.Next
		tmpHead = tmpHead.Next
		index++
	}

	return res
}
