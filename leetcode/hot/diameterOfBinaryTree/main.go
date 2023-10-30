package main

// https://leetcode.cn/problems/diameter-of-binary-tree/?envType=study-plan-v2&envId=top-100-liked
// 你的审题上还是有问题的

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type info struct {
	height  int
	maxPath int
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxDiameter(root).maxPath - 1
}

func maxDiameter(root *TreeNode) *info {
	if root == nil {
		return &info{
			height:  0,
			maxPath: 0,
		}
	}
	rootInfo := new(info)
	leftInfo := maxDiameter(root.Left)
	rightInfo := maxDiameter(root.Right)
	rootInfo.height = max(leftInfo.height, rightInfo.height) + 1
	rootInfo.maxPath = max(max(leftInfo.maxPath, rightInfo.maxPath), leftInfo.height+rightInfo.height+1)
	return rootInfo
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
