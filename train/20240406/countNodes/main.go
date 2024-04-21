package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count func(root *TreeNode) int
	count = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return count(root.Right) + count(root.Left) + 1
	}

	return count(root)

}

func maxV(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1

}
