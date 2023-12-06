package main

import "fmt"

func main() {
	root := new(TreeNode)
	root.Val = 2
	root.Left = &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Right = &TreeNode{
		Val: 3,
	}
	printTree(root)
	flatten(root)
	fmt.Println("-------")
	printTree(root)
}

func printTree(root *TreeNode) {
	for root == nil {
		fmt.Println("null")
		return
	}

	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 时间和空间复杂度都不太行
//func flatten(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	res := make([]int, 0)
//	process(root, &res)
//	root.Left = nil
//	resRoot := root
//	for i := 1; i < len(res); i++ {
//		resRoot.Right = new(TreeNode)
//		resRoot.Right.Val = res[i]
//		resRoot = resRoot.Right
//	}
//}
//
//func process(root *TreeNode, res *[]int) {
//	if root == nil {
//		return
//	}
//	*res = append(*res, root.Val)
//	process(root.Left, res)
//	process(root.Right, res)
//}

// 问题应该出在指针并没有变换，但是不知道怎么解决
//func flatten(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	head := new(TreeNode)
//	head.Right = new(TreeNode)
//	tmp := head
//	process(root, tmp)
//	root = head.Right
//}
//
//func process(root, pre *TreeNode) {
//	if root == nil {
//		return
//	}
//	pre.Right = new(TreeNode)
//	pre.Right.Val = root.Val
//	pre = pre.Right
//	process(root.Left, pre)
//	process(root.Right, pre)
//}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	tmp := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmp
}
