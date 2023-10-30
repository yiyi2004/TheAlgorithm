package main

func main() {

}

// https://leetcode.cn/problems/symmetric-tree/description/?envType=study-plan-v2&envId=top-100-liked

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 第一次的错解
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(left *TreeNode, right *TreeNode) bool {
	// 这个是终止条件，因为是递归往上传递的，有一个 false，整体都是 false
	// 所以一条路径如果是 true，那么一定会到达最下面
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
}
