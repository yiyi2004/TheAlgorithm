package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	return process(root, res)
}

func process(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}

	res = process(root.Left, res)
	res = append(res, root.Val)
	res = process(root.Right, res)
	return res
}
