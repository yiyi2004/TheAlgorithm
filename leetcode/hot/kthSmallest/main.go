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
	res := process(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	return process(root)[k-1]
}

func process(root *TreeNode) []int {
	res := make([]int, 0)
	inOrder(root, &res)
	return res
}

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

// 中序遍历应该是可以解出来的，但是代码有点不会写了
//var index = 0

//func kthSmallest(root *TreeNode, k int) int {
//	var val *int
//	*val = math.MinInt
//	process(root, index, k, val)
//	return *val
//}

//func process(root *TreeNode, index, k int, res *int) {
//	if root == nil {
//		return
//	}
//	index++
//	process(root.Left, index, k, res)
//	if index == k {
//		*res = root.Val
//		return
//	}
//	index++
//	process(root.Right, index, k, res)
//}
