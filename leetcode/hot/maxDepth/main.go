package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return process(root)
}

func process(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := process(root.Left)
	rightHeight := process(root.Right)
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}
